package neovim

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/neovim/go-client/nvim"
)

type KeyMaps struct {
	api    *Api
	set    func(mode Mode, lhs string, rhs string, opts map[string]bool) error
	get    func(mode Mode) ([]*nvim.Mapping, error)
	delete func(mode Mode, lhs string)
}

func newBufferKeyMaps(api *Api, id nvim.Buffer) KeyMaps {
	return KeyMaps{
		api: api,
		set: func(mode Mode, lhs string, rhs string, opts map[string]bool) error {
			return api.nvim().SetBufferKeyMap(id, string(mode), lhs, rhs, opts)
		},
		get: func(mode Mode) ([]*nvim.Mapping, error) {
			return api.nvim().BufferKeyMap(id, string(mode))
		},
		delete: func(mode Mode, lhs string) {
			api.nvim().DeleteBufferKeyMap(id, string(mode), lhs)
		},
	}
}

func newGlobalKeyMaps(api *Api) KeyMaps {
	return KeyMaps{
		api: api,
		set: func(mode Mode, lhs string, rhs string, opts map[string]bool) error {
			return api.nvim().SetKeyMap(string(mode), lhs, rhs, opts)
		},
		get: func(mode Mode) ([]*nvim.Mapping, error) {
			return api.nvim().KeyMap(string(mode))
		},
		delete: func(mode Mode, lhs string) {
			api.nvim().DeleteKeyMap(string(mode), lhs)
		},
	}
}

func (m *KeyMaps) SetFunc(mode Mode, keys string, fn func()) {
	handler := m.api.Handler.Create(fn)
	eval := fmt.Sprintf(`:silent call %s<CR>`, handler)
	m.set(mode, keys, eval, map[string]bool{"silent": true, "nowait": true})
}

func (m *KeyMaps) SetTextAction(keys string, fn func(string) string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("error: %v", err)
		}
	}()

	h := m.textActionHandler(fn)

	m.Setf(ModeNormal, keys, `:<C-U>silent call %s<CR>g@`, m.createMotionActionHandler(h))
	m.Setf(Mode("x"), keys, `:<C-U>silent call %s<CR>`, m.createSelectionActionHandler(h))
}

func (m KeyMaps) Set(mode Mode, keys string, eval string) {
	m.set(mode, keys, eval, map[string]bool{"silent": true, "nowait": true})
}

func (m KeyMaps) Setf(mode Mode, keys string, eval string, args ...interface{}) {
	m.set(mode, keys, fmt.Sprintf(eval, args...), map[string]bool{"silent": true, "nowait": true})
}

func (m KeyMaps) Delete(mode Mode, keys string) {
	m.delete(mode, keys)
}

func (m KeyMaps) Disable(mode Mode, keys string) {
	m.set(mode, keys, "<nop>", map[string]bool{"silent": true, "nowait": true})
}

func (m *KeyMaps) createSelectionActionHandler(fn func(...interface{})) string {
	return m.api.Handler.Create(fn).StringWithEvals("visualmode()")
}

func (m *KeyMaps) createMotionActionHandler(fn func(...interface{})) string {
	return m.api.Handler.Create(func() {
		m.api.Handler.SetOperatorFunc(func(args []interface{}) {
			fn(args...)
		})
	}).String()
}

// Taken from https://github.com/arthurxavierx/vim-caser/blob/master/plugin/caser.vim#L17
func (m KeyMaps) textActionHandler(fn func(string) string) func(typi ...interface{}) {

	return func(typi ...interface{}) {

		// backup settings that we will change
		opt := m.api.Global.Options
		selection := opt.Selection()
		clipboard := opt.Clipboard()

		// make selection and clipboard work the way we need
		opt.SetSelection(GlobalSelectionInclusive)
		opt.SetClipboard(GlobalClipboardNone)

		// backup the unnamed register, which we will be yanking into
		regBak, _ := m.api.Execute("echo @@")

		// restore saved settings and register value
		defer func() {
			opt.SetSelection(selection)
			opt.SetClipboard(clipboard)

			m.api.Executef("let @@ = '%s'", strings.Replace(regBak, "'", `'."'".'`, -1))
		}()

		//

		typ := ""
		if len(typi) != 1 {
			return
		}

		typ, _ = typi[0].(string)

		sel := ""
		if regexp.MustCompile(`^\d+$`).MatchString(typ) {
			// if type is a number, then select that many lines
			sel = fmt.Sprintf("V%s$y", typ)

		} else if len(typ) == 1 {
			// if type is 'v', 'V', or '<C-V>' (i.e. 0x16) then reselect the visual region
			sel = fmt.Sprintf("`<%s`>y", typ)

		} else if typ == "line" {
			// line-based text motion
			sel = "'[V']y"

		} else if typ == "line" {
			// block-based text motion
			sel = "`[\\<C-V>`]y"

		} else { // == "char"
			// char-based text motion
			sel = "`[v`]y"
		}

		m.api.Executef("normal! %s", sel)
		content, _ := m.api.Execute("echo @@")

		replacement := fn(content)

		// put the replacement text into the unnamed register, and also set it to be a
		// characterwise, linewise, or blockwise selection, based upon the selection type of the
		// yank we did above
		m.api.Executef("call setreg('@', '%s', getregtype('@'))", strings.Replace(replacement, "'", `'."'".'`, -1))

		// reselect the visual region and paste
		m.api.Executef("normal! gvp")
	}
}
