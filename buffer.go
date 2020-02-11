package neovim

import (
	"fmt"
	"strings"
	"sync"

	"github.com/josa42/go-neovim/disposables"
	"github.com/neovim/go-client/nvim"
)

type Buffer struct {
	api         *Api
	id          nvim.Buffer
	Options     BufferOptions
	KeyMaps     KeyMaps
	Vars        Vars
	mutex       *sync.Mutex
	disposables *disposables.Collection
}

func newBufferById(api *Api, id nvim.Buffer) *Buffer {

	if r := api.registry.get(RegistryTypeBuffer, int(id)); r != nil {
		if b, ok := r.(*Buffer); ok {
			return b
		}
	}

	b := &Buffer{
		api:         api,
		id:          nvim.Buffer(id),
		Options:     BufferOptions{api: api, bufferID: id},
		KeyMaps:     newBufferKeyMaps(api, id),
		Vars:        newBufferVars(api, id),
		mutex:       &sync.Mutex{},
		disposables: disposables.NewCollection(),
	}

	api.registry.add(RegistryTypeBuffer, b)

	return b
}

func (b *Buffer) ID() int {
	return int(b.id)
}

func (b *Buffer) Exists() bool {
	bs, _ := b.api.nvim().Buffers()
	for _, bi := range bs {
		if bi == b.id {
			return true
		}
	}

	return false
}

func (b *Buffer) IsCurrent() bool {
	currentID, _ := b.api.nvim().CurrentBuffer()
	return currentID > 0 && b.id == currentID
}

func (b *Buffer) Close() {
	b.api.nvim().DetachBuffer(nvim.Buffer(b.id))
	b.api.Executef("bwipeout %d", b.id)
	b.disposables.Dispose()
}

func (b *Buffer) Path() string {
	var path string
	b.api.nvim().Call("expand", &path, fmt.Sprintf(`#%d:p`, b.id))
	return path
}

func (b *Buffer) Windows() []*Window {
	bufferWindows := []*Window{}

	wins, _ := b.api.nvim().Windows()
	for _, win := range wins {
		wb, _ := b.api.nvim().WindowBuffer(win)
		if wb == b.id {
			if bwin, found := b.api.WindowById(int(win)); found {
				bufferWindows = append(bufferWindows, bwin)
			}
		}
	}

	return bufferWindows
}

////////////////////////////////////////////////////////////////////////////////
// Options

func (b *Buffer) Freeze() {
	b.Options.SetModifiable(false)
	b.Options.SetReadOnly(true)
}

func (b *Buffer) Unfreeze() {
	b.Options.SetModifiable(true)
	b.Options.SetReadOnly(false)
}

func (b *Buffer) Title() string {
	title, _ := b.api.nvim().BufferName(nvim.Buffer(b.id))
	return title
}
func (b *Buffer) SetTitle(title string) {
	b.api.nvim().SetBufferName(nvim.Buffer(b.id), title)
}

////////////////////////////////////////////////////////////////////////////////
// Vars

func (b *Buffer) VarString(name string) string {
	var value string
	b.api.nvim().BufferVar(b.id, name, &value)
	return value
}

func (b *Buffer) SetVarString(name string, value string) {
	b.api.nvim().SetBufferVar(b.id, name, value)
}

func (b *Buffer) VarBool(name string) bool {
	var value bool
	b.api.nvim().BufferVar(b.id, name, &value)
	return value
}

func (b *Buffer) SetVarBool(name string, value bool) {
	b.api.nvim().SetBufferVar(b.id, name, value)
}

////////////////////////////////////////////////////////////////////////////////
// Content

func (b *Buffer) Lines() []string {
	bs, _ := b.api.nvim().BufferLines(b.id, 0, -1, false)

	lines := []string{}
	for _, b := range bs {
		lines = append(lines, string(b))
	}
	return lines
}

func (b *Buffer) SetLines(lines []string) {
	defer b.lock()()
	defer b.makeWritable()()

	bb := [][]byte{}
	for _, l := range lines {
		bb = append(bb, []byte(l))
	}

	batch := b.api.nvim().NewBatch()
	batch.SetBufferLines(b.id, 0, -1, false, bb)
	batch.Execute()

	// TODO debug why diff does not work correctly
	// currentLines := b.Lines()
	// currentCount := len(currentLines)
	//
	// batch := b.api.Nvim().NewBatch()
	//
	// for lineIdx, l := range lines {
	// 	cl := ""
	// 	if lineIdx < currentCount {
	// 		cl = currentLines[lineIdx]
	// 	}
	//
	// 	if l != cl {
	// 		batch.SetBufferLines(b.id, lineIdx, lineIdx, true, [][]byte{[]byte(l)})
	// 		log.Printf("%02d: %s", lineIdx, l)
	// 	}
	// }
	//
	// batch.SetBufferLines(b.id, len(lines), -1, false, [][]byte{})
	// batch.Execute()
}

func (b *Buffer) IsEmpty() bool {
	n := b.api.nvim()

	if c, _ := n.BufferLineCount(b.id); c > 1 {
		return false
	}

	if lines, _ := n.BufferLines(b.id, 0, -1, false); string(lines[0]) != "" {
		return false
	}
	return true
}

func (b *Buffer) makeWritable() func() {
	readonly := b.Options.ReadOnly()
	modifiable := b.Options.Modifiable()

	if readonly {
		b.Options.SetReadOnly(false)
	}

	if !modifiable {
		b.Options.SetModifiable(true)
	}

	return func() {
		if readonly {
			b.Options.SetReadOnly(readonly)
		}
		if !modifiable {
			b.Options.SetModifiable(modifiable)
		}
	}
}

func (b *Buffer) lock() func() {
	b.mutex.Lock()
	return func() {
		b.mutex.Unlock()
	}
}

////////////////////////////////////////////////////////////////////////////////

func (b *Buffer) On(event string, fn func()) {
	handler := b.api.Handler.Create(fn)
	b.disposables.Add(handler)

	groupName := fmt.Sprintf("buffer_%s", handler.uuid)

	b.api.Executef("augroup %s | autocmd %s <buffer=%d> call %s | augroup END", groupName, event, b.ID(), handler)
	b.disposables.Add(disposables.New(func() {
		b.api.Executef("autocmd! %s", groupName)
	}))
}

////////////////////////////////////////////////////////////////////////////////
// API

func (api *Api) BufferById(id int) (*Buffer, bool) {
	buffer := newBufferById(api, nvim.Buffer(id))
	return buffer, buffer.Exists()
}

func (api *Api) CurrentBuffer() *Buffer {
	id, _ := api.nvim().CurrentBuffer()
	return newBufferById(api, nvim.Buffer(id))
}

func (api *Api) FindBuffer(fn func(buffer *Buffer) bool) (*Buffer, bool) {
	bufferIDs, _ := api.p.Nvim.Buffers()
	for _, id := range bufferIDs {
		buffer := newBufferById(api, id)
		if fn(buffer) {
			return buffer, true
		}
	}

	return newBufferById(api, 0), false
}

type SplitModifier int

const (
	SplitVertical SplitModifier = iota
	SplitHorizontal
	SplitTopLeft
	SplitBottomRight
)

func (s SplitModifier) String() string {
	switch s {
	case SplitVertical:
		return "vertical"
	case SplitHorizontal:
		return "horizontal"
	case SplitTopLeft:
		return "topleft"
	case SplitBottomRight:
		return "botright"
	default:
		return ""
	}
}

func (api *Api) CreateSplitBuffer(width int, mods ...SplitModifier) *Buffer {
	var bID nvim.Buffer

	modStrs := []string{}
	for _, m := range mods {
		modStrs = append(modStrs, m.String())
	}

	batch := api.nvim().NewBatch()
	batch.Command(fmt.Sprintf("%s %d new", strings.Join(modStrs, " "), width))
	batch.CurrentBuffer(&bID)
	batch.Execute()

	buffer, _ := api.BufferById(int(bID))

	return buffer
}
