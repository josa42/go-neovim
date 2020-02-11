package neovim

import (
	"github.com/neovim/go-client/nvim"
)

type Vars struct {
	get    func(name string, result interface{}) error
	set    func(name string, value interface{}) error
	delete func(name string) error
}

func newBufferVars(api *Api, id nvim.Buffer) Vars {
	return Vars{
		get: func(name string, result interface{}) error {
			return api.nvim().BufferVar(id, name, result)
		},
		set: func(name string, value interface{}) error {
			return api.nvim().SetBufferVar(id, name, value)
		},
		delete: func(name string) error {
			return api.nvim().DeleteBufferVar(id, name)
		},
	}
}

func newWindowVars(api *Api, id nvim.Window) Vars {
	return Vars{
		get: func(name string, result interface{}) error {
			return api.nvim().WindowVar(id, name, result)
		},
		set: func(name string, value interface{}) error {
			return api.nvim().SetWindowVar(id, name, value)
		},
		delete: func(name string) error {
			return api.nvim().DeleteWindowVar(id, name)
		},
	}
}

func newTabVars(api *Api, id nvim.Tabpage) Vars {
	return Vars{
		get: func(name string, result interface{}) error {
			return api.nvim().TabpageVar(id, name, result)
		},
		set: func(name string, value interface{}) error {
			return api.nvim().SetTabpageVar(id, name, value)
		},
		delete: func(name string) error {
			return api.nvim().DeleteTabpageVar(id, name)
		},
	}
}

func newGlobalVars(api *Api) Vars {
	return Vars{
		get: func(name string, result interface{}) error {
			return api.nvim().Var(name, result)
		},
		set: func(name string, value interface{}) error {
			return api.nvim().SetVar(name, value)
		},
		delete: func(name string) error {
			return api.nvim().DeleteVar(name)
		},
	}
}

func (v *Vars) String(name string) string {
	var value string
	v.get(name, &value)
	return value
}

func (v *Vars) SetString(name string, value string) {
	v.set(name, value)
}

func (v *Vars) Bool(name string) bool {
	var value bool
	v.get(name, &value)
	return value
}

func (v *Vars) SetBool(name string, value bool) {
	v.set(name, value)
}

func (v *Vars) Int(name string) int {
	var value int
	v.get(name, &value)
	return value
}

func (v *Vars) SetInt(name string, value int) {
	v.set(name, value)
}

func (v *Vars) Delete(name string) {
	v.delete(name)
}

