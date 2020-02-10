package neovim

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/josa42/go-neovim/disposables"
	"github.com/kardianos/osext"
)

type HandlerFunc struct {
	uuid         string
	functionName string
	disposeFn    func()
}

func (h *HandlerFunc) String() string {
	return h.StringWithEvals()
}

func (h *HandlerFunc) StringWithEvals(evals ...string) string {
	args := append([]string{fmt.Sprintf(`'%s'`, h.uuid)}, evals...)
	return fmt.Sprintf(`%s(%s)`, h.functionName, strings.Join(args, ", "))
}

func (h *HandlerFunc) Dispose() {
	h.disposeFn()
}

type Handler struct {
	api          *Api
	uuid         string
	handlers     map[string]func([]interface{})
	operatorFunc func(args []interface{})
}

func newHandler(api *Api) Handler {
	h := Handler{
		api:      api,
		uuid:     readHandlerUUID(),
		handlers: map[string]func([]interface{}){},
	}

	// go func() {
	// 	for true {
	// 		l := []string{}
	// 		l = append(l, fmt.Sprintf("Handlers: %d", len(h.handlers)))
	// 		for k, _ := range h.handlers {
	// 			l = append(l, fmt.Sprintf("- %s", k))
	// 		}
	// 		log.Print(strings.Join(l, "\n"))
	//
	// 		time.Sleep(2 * time.Second)
	// 	}
	// }()

	return h
}

func (h *Handler) register(api *Api) {
	api.Function(h.functionName(), func(args []interface{}) error {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("GlobalHandler() recover: %v\n", err)
			}
		}()
		if len(args) > 0 {
			if hID, ok := args[0].(string); ok {
				if hndl, ok := h.handlers[hID]; ok {
					hndl(args[1:])
					return nil
				}
			}
		}
		return nil
	})

	api.Function(h.operatorFunctionName(), func(args []interface{}) error {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("OperatorFunc() recover: %v\n", err)
			}
		}()

		if h.operatorFunc != nil {
			h.operatorFunc(args)
		}
		return nil
	})
}

func (h *Handler) Create(fn interface{}) *HandlerFunc {
	uuid := generateUUID()

	if fnh, ok := fn.(func()); ok {
		h.handlers[uuid] = func([]interface{}) {
			fnh()
		}
	} else if fnh, ok := fn.(func(args ...interface{})); ok {
		h.handlers[uuid] = func(args []interface{}) {
			fnh(args...)
		}
	} else {
		panic("invalid handler")
	}

	return &HandlerFunc{
		uuid:         uuid,
		functionName: h.functionName(),
		disposeFn: func() {
			delete(h.handlers, uuid)
		},
	}
}

func (h *Handler) SetOperatorFunc(fn func(args []interface{})) disposables.Disposable {
	h.operatorFunc = fn
	h.api.Global.Options.SetOperatorFunc(h.operatorFunctionName())

	return disposables.New(func() {
		h.api.Global.Options.SetOperatorFunc("")
		h.operatorFunc = nil
	})
}

func (h *Handler) operatorFunctionName() string {
	return fmt.Sprintf(`OperatorFunc_%s`, h.uuid)
}

func (h *Handler) functionName() string {
	return fmt.Sprintf(`Handler_%s`, h.uuid)
}

// TODO Find a better way to generate and store uuid
func readHandlerUUID() string {
	executable, _ := osext.Executable()
	uuidFile := fmt.Sprintf("%s.uuid", executable)

	content, _ := ioutil.ReadFile(uuidFile)
	uuid := string(content)

	if uuid == "" {
		b := make([]byte, 16)
		rand.Read(b)
		uuid = generateUUID()

		ioutil.WriteFile(uuidFile, []byte(uuid), 0777)
	}

	return uuid
}

func generateUUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
