package neovim

import "github.com/neovim/go-client/nvim"

type Tab struct {
	api  *Api
	id   nvim.Tabpage
	Vars Vars
}

func newTabById(api *Api, id nvim.Tabpage) *Tab {
	if r := api.registry.get(RegistryTypeTab, int(id)); r != nil {
		if tab, ok := r.(*Tab); ok {
			return tab
		}
	}

	tab := &Tab{
		api:  api,
		id:   nvim.Tabpage(id),
		Vars: newTabVars(api, id),
	}

	api.registry.add(RegistryTypeTab, tab)

	return tab
}

func (t *Tab) ID() int {
	return int(t.id)
}

func (tab *Tab) Exists() bool {
	ts, _ := tab.api.nvim().Tabpages()
	for _, t := range ts {
		if t == tab.id {
			return true
		}
	}

	return false
}

func (b *Tab) IsCurrent() bool {
	currentID, _ := b.api.nvim().CurrentTabpage()
	return currentID > 0 && b.id == currentID
}

func (t *Tab) HasBufferID(bID int) bool {
	wins, _ := t.api.nvim().TabpageWindows(nvim.Tabpage(t.ID()))

	for _, win := range wins {
		wb, _ := t.api.nvim().WindowBuffer(win)
		if int(wb) == bID {
			return true
		}
	}

	return false
}

func (t *Tab) HasBuffer(b Buffer) bool {
	return t.HasBufferID(b.ID())
}

func (t *Tab) Windows() []*Window {
	wins, _ := t.api.nvim().TabpageWindows(nvim.Tabpage(t.ID()))

	windows := []*Window{}
	for _, winID := range wins {
		windows = append(windows, newWindowById(t.api, winID))
	}

	return windows
}

func (t *Tab) FindWindow(fn func(window *Window) bool) (*Window, bool) {
	for _, win := range t.Windows() {
		if fn(win) {
			return win, true
		}
	}

	return newWindowById(t.api, 0), false
}

////////////////////////////////////////////////////////////////////////////////
// LIfe Cycle

func (t *Tab) Close(force bool) {
	for _, win := range t.Windows() {
		win.Close(force)
	}
}

////////////////////////////////////////////////////////////////////////////////

func (api *Api) TabById(id int) (*Tab, bool) {
	buffer := newTabById(api, nvim.Tabpage(id))
	return buffer, buffer.Exists()
}

func (api *Api) CurrentTab() *Tab {
	id, _ := api.nvim().CurrentTabpage()
	return newTabById(api, nvim.Tabpage(id))
}

func (api *Api) FindTab(fn func(tab *Tab) bool) (*Tab, bool) {
	bufferIDs, _ := api.nvim().Tabpages()
	for _, id := range bufferIDs {
		buffer := newTabById(api, id)
		if fn(buffer) {
			return buffer, true
		}
	}

	return newTabById(api, 0), false
}
