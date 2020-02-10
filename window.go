package neovim

import "github.com/neovim/go-client/nvim"

type Window struct {
	api     *Api
	id      nvim.Window
	Vars    Vars
	Options WindowOptions
}

func newWindowById(api *Api, id nvim.Window) *Window {
	if r := api.registry.get(RegistryTypeWindow, int(id)); r != nil {
		if win, ok := r.(*Window); ok {
			return win
		}
	}

	win := &Window{
		api:     api,
		id:      nvim.Window(id),
		Vars:    newWindowVars(api, id),
		Options: WindowOptions{api: api, windowID: id},
	}

	api.registry.add(RegistryTypeWindow, win)

	return win
}

func (w *Window) ID() int {
	return int(w.id)
}

func (win *Window) Exists() bool {
	ws, _ := win.api.nvim().Windows()
	for _, w := range ws {
		if w == win.id {
			return true
		}
	}

	return false
}

func (b *Window) IsCurrent() bool {
	currentID, _ := b.api.nvim().CurrentWindow()
	return currentID > 0 && b.id == currentID
}

func (w *Window) Focus() {
	w.api.nvim().SetCurrentWindow(nvim.Window(w.id))
}

func (w *Window) Buffer() *Buffer {
	bID, _ := w.api.nvim().WindowBuffer(nvim.Window(w.id))
	return newBufferById(w.api, bID)
}

func (w *Window) Tab() *Tab {
	tID, _ := w.api.nvim().WindowTabpage(nvim.Window(w.id))
	return newTabById(w.api, tID)
}

func (w *Window) Cursor() Cursor {
	pos, _ := w.api.nvim().WindowCursor(nvim.Window(w.id))
	return Cursor(pos)
}

func (w *Window) SetCursor(c Cursor) {
	w.api.nvim().SetWindowCursor(nvim.Window(w.id), c)
}

////////////////////////////////////////////////////////////////////////////////
// Life Cycle

func (w *Window) Close(force bool) {
	w.api.nvim().CloseWindow(nvim.Window(w.id), force)
}

// Vars

func (b *Window) VarString(name string) string {
	var value string
	b.api.nvim().WindowVar(b.id, name, &value)
	return value
}

func (b *Window) SetVarString(name string, value string) {
	b.api.nvim().SetWindowVar(b.id, name, value)
}

func (b *Window) VarBool(name string) bool {
	var value bool
	b.api.nvim().WindowVar(b.id, name, &value)
	return value
}

func (b *Window) SetVarBool(name string, value bool) {
	b.api.nvim().SetWindowVar(b.id, name, value)
}

////////////////////////////////////////////////////////////////////////////////
// API

func (api *Api) WindowById(id int) (*Window, bool) {
	win := newWindowById(api, nvim.Window(id))
	return win, win.Exists()
}

func (api *Api) CurrentWindow() *Window {
	id, _ := api.nvim().CurrentWindow()
	return newWindowById(api, nvim.Window(id))
}

func (api *Api) FindWindow(fn func(win *Window) bool) (*Window, bool) {
	bufferIDs, _ := api.p.Nvim.Windows()
	for _, id := range bufferIDs {
		win := newWindowById(api, id)
		if fn(win) {
			return win, true
		}
	}

	return newWindowById(api, 0), false
}
