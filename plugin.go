package neovim

import (
	"time"

	"github.com/neovim/go-client/nvim/plugin"
)

type Plugin interface {
	Register(api RegisterApi)
	Activate(api *Api)
}

type Registerable interface {
	Register(api RegisterApi)
}

func Register(p Plugin) {

	plugin.Main(func(np *plugin.Plugin) error {
		api := newApiWithPlugin(np)
		api.Handler.register(api)

		if p, ok := p.(Registerable); ok {
			p.Register(api)
		}

		// TODO find a better solution to call the Axctivate hook
		go func() {
			time.Sleep(10 * time.Millisecond)
			p.Activate(api)
		}()

		return nil
	})
}
