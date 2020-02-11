package neovim

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/neovim/go-client/nvim"
	"github.com/neovim/go-client/nvim/plugin"
)

type RegisterApi interface {
	Function(name string, fn interface{})
}

type Api struct {
	p *plugin.Plugin

	Out      Out
	Global   Global
	Handler  Handler
	registry *registry
	Renderer Renderer
}

func newApiWithPlugin(p *plugin.Plugin) *Api {
	api := &Api{p: p}
	api.Out = Out{api: api}
	api.Global = newGlobal(api)
	api.Handler = newHandler(api)
	api.registry = newRegistry(api)
	api.Renderer = Renderer{}

	return api
}

func NewApi() *Api {
	stdout := os.Stdout
	os.Stdout = os.Stderr

	v, err := nvim.New(os.Stdin, stdout, stdout, log.Printf)
	if err != nil {
		log.Fatal(err)
	}

	return newApiWithPlugin(plugin.New(v))
}

func (api *Api) nvim() *nvim.Nvim {
	return api.p.Nvim
}

// }

func (api *Api) Execute(cmd string) (string, error) {
	return api.nvim().CommandOutput(cmd)
}

func (api *Api) Executef(format string, args ...interface{}) (string, error) {
	return api.Execute(fmt.Sprintf(format, args...))
}

func (api *Api) Function(name string, fn interface{}) {
	api.p.HandleFunction(&plugin.FunctionOptions{Name: name}, fn)
}

func (api *Api) Cwd() string {
	var cwd string
	api.nvim().Call("getcwd", &cwd)
	return cwd
}

func (api *Api) on(event, pattern string, fn func()) {
	api.p.HandleAutocmd(&plugin.AutocmdOptions{Event: event, Pattern: pattern}, wrapEventHandler(fn))
}

func wrapEventHandler(fn func()) interface{} {
	return func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("handler recover (1): %v\n", err)
				log.Println(string(debug.Stack()))
			}
		}()
		fn()
	}
}
