package view

import (
	"fmt"

	"github.com/josa42/go-neovim"
	"github.com/josa42/go-neovim/disposables"
)

const (
	ItemStatusChanged    = '◎'
	ItemStatusAdded      = '⦿'
	ItemStatusConflicted = '◉'
)

const (
	LevelPrefix = "  "
)

type TreeProvider interface {
	FileType() string
	Root() TreeItem
}

type TreeItem interface {
	String() string
	Children() []TreeItem
}

type Openable interface {
	IsOpenable() bool
	IsOpen() bool
	Open()
	Close()
}

type Statusable interface {
	Status() rune
}

type TreeAction struct {
	Mode    string
	Keys    string
	Handler func(TreeItem)
}

type ActionableTree interface {
	Actions() []TreeAction
}

type Changable interface {
	Listen(func())
	Unlisten()
}

type line struct {
	prefix string
	item   TreeItem
}

func (l *line) String() string {
	status := ' '
	if i, ok := l.item.(Statusable); ok {
		status = i.Status()
	}

	return fmt.Sprintf("%s%s %s", l.prefix, string(status), l.item.String())
}

// Interface Assertions
var _ neovim.View = (*TreeView)(nil)
var _ neovim.Initializable = (*TreeView)(nil)
var _ disposables.Disposable = (*TreeView)(nil)

type TreeView struct {
	renderer    neovim.ViewRenderer
	provider    TreeProvider
	lines       []line
	disposables *disposables.Collection
}

func NewTreeView(provider TreeProvider) *TreeView {
	return &TreeView{
		provider:    provider,
		disposables: disposables.NewCollection(),
	}
}
func (t *TreeView) FileType() string {
	return t.provider.FileType()
}

func (t *TreeView) Attach(r neovim.ViewRenderer) {
	t.renderer = r
}

func (t *TreeView) Rerender() {
	t.renderer.ShouldRender()
}

func (t *TreeView) Initialize(b *neovim.Buffer, api *neovim.Api) {

	b.On(neovim.EventCursorMoved, func() {
		w := api.CurrentWindow()
		y := w.Cursor().Y()
		w.SetCursor(neovim.Cursor{y, 0})
		api.Execute("normal! <C-c>")
		api.Execute("set nohlsearch")

	})

	nopKeyMaps := []string{"i", "a", "v", "V", "<C>", "<Leader>", "<C-v>", "<C-0>", "h", "l", "<Left>", "<Right>", "0", "$", "^"}

	for _, m := range nopKeyMaps {
		b.KeyMaps.Disable(neovim.ModeAll, m)
	}

	if p, ok := t.provider.(ActionableTree); ok {
		for _, a := range p.Actions() {
			func(a TreeAction) {
				fn := api.Handler.Create(func() {
					win := api.CurrentWindow()
					cursor := win.Cursor()
					idx := cursor.Y() - 1

					if idx >= 0 && idx < len(t.lines) {
						a.Handler(t.lines[idx].item)
						t.renderer.ShouldRender()
					}
				})
				t.disposables.Add(fn)
				b.KeyMaps.Set(neovim.ModeNormal, a.Keys, fmt.Sprintf(`:silent call %s<CR>`, fn))
			}(a)
		}
	}

	if p, ok := t.provider.(Changable); ok {
		p.Listen(func() {
			t.renderer.ShouldRender()
		})
	}
}

func (t *TreeView) Dispose() {
	if p, ok := t.provider.(Changable); ok {
		p.Unlisten()
	}
	t.disposables.Dispose()
	t.disposables = disposables.NewCollection()
}

func (t *TreeView) Update() {
	if p, ok := t.provider.(neovim.Updatable); ok {
		p.Update()
	}

	t.lines = t.visibleLines("", t.provider.Root().Children())
}

func (t *TreeView) visibleLines(prefix string, items []TreeItem) []line {
	lines := []line{}

	for _, i := range items {
		lines = append(lines, line{prefix: prefix, item: i})
		if t.shouldRenderChildren(i) {
			lines = append(lines, t.visibleLines(prefix+LevelPrefix, i.Children())...)
		}
	}

	return lines
}

func (t *TreeView) shouldRenderChildren(i TreeItem) bool {

	if item, ok := i.(Openable); ok {
		return item.IsOpen()
	}

	return true
}

func (t *TreeView) Lines() []string {
	lines := []string{}

	for _, l := range t.lines {
		lines = append(lines, l.String())
	}

	return lines
}
