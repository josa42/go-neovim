package main

import (
	"strings"

	"github.com/josa42/go-neovim"
)

func main() {
	neovim.Register(&plugin{})
}

type plugin struct{}

func (p *plugin) Activate(api *neovim.Api) {

	api.Global.KeyMaps.SetTextAction("cu", func(s string) string {
		return strings.ToUpper(s)
	})

	api.Global.KeyMaps.SetTextAction("cl", func(s string) string {
		return strings.ToLower(s)
	})
}

