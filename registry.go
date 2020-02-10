package neovim

import (
	"log"
	"sync"
)

const (
	RegistryTypeBuffer = "buffer"
	RegistryTypeWindow = "window"
	RegistryTypeTab    = "tab"
)

type registerable interface {
	ID() int
	Exists() bool
}

type registry struct {
	content map[string][]registerable
	mutex   sync.Mutex
	api     *Api
}

func newRegistry(api *Api) *registry {
	r := &registry{
		content: map[string][]registerable{},
		mutex:   sync.Mutex{},
		api:     api,
	}

	api.On(EventBufDelete, "*", func() {
		r.garbadgeCollect(RegistryTypeBuffer)
	})

	api.On(EventBufWipeout, "*", func() {
		r.garbadgeCollect(RegistryTypeBuffer)
	})

	api.On(EventTabClosed, "*", func() {
		r.garbadgeCollect(RegistryTypeTab)
	})

	api.On(EventBufWinLeave, "*", func() {
		r.garbadgeCollect(RegistryTypeWindow)
	})

	// go func() {
	// 	for true {
	// 		l := []string{}
	// 		l = append(l, fmt.Sprintf("Registry: %d", len(r.content)))
	// 		for t, c := range r.content {
	// 			l = append(l, fmt.Sprintf("- %s [%d]", t, len(c)))
	// 		}
	// 		log.Print(strings.Join(l, "\n"))
	//
	// 		time.Sleep(2 * time.Second)
	// 	}
	// }()

	return r
}

func (r *registry) lock() func() {
	r.mutex.Lock()
	return func() {
		r.mutex.Unlock()
	}
}

func (r *registry) garbadgeCollect(t string) {
	defer r.lock()()

	if c, ok := r.content[t]; ok {
		content := []registerable{}
		for _, i := range c {
			if i.Exists() {
				content = append(content, i)
			} else {
				log.Printf("Remove %s %d", t, i.ID())
			}
		}
		r.content[t] = content
	}
}

func (r *registry) add(t string, i registerable) {
	if _, ok := r.content[t]; !ok {
		r.content[t] = []registerable{}
	}

	r.content[t] = append(r.content[t], i)
}

func (r *registry) get(t string, idx int) registerable {
	if c, ok := r.content[t]; ok {
		for _, i := range c {
			if i.ID() == idx {
				return i
			}
		}
	}

	return nil
}
