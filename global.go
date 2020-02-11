package neovim

import (
	"fmt"
	"log"
)

type Global struct {
	api     *Api
	Vars    Vars
	Options GlobalOptions
	KeyMaps KeyMaps
}

func newGlobal(api *Api) Global {
	return Global{
		api:     api,
		Vars:    newGlobalVars(api),
		Options: GlobalOptions{api},
		KeyMaps: newGlobalKeyMaps(api),
	}
}

func (g *Global) On(event string, fn func()) {
	handler := g.api.Handler.Create(fn)
	groupName := fmt.Sprintf("global_%s", handler.uuid)
	log.Printf("augroup %s | autocmd %s * call %s | augroup END", groupName, event, handler)
	g.api.Executef("augroup %s | autocmd %s * call %s | augroup END", groupName, event, handler)
}

