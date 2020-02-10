package neovim

import (
	"log"
	"time"

	"github.com/josa42/go-neovim/disposables"
)

type View interface {
	FileType() string
	Attach(ViewRenderer)
	Lines() []string
}

type Initializable interface {
	Initialize(*Buffer, *Api)
}

type Updatable interface {
	Update()
}

type ViewRenderer struct {
	buffer *Buffer
	view   View
}

func (r *ViewRenderer) ShouldRender() {
	r.render()
}

func (r *ViewRenderer) render() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Renderer:render(): Recover: %v\n", err)
		}
	}()

	if r.view == nil {
		log.Println("Renderer:render(): View is nil!?")
		return
	}

	if v, ok := r.view.(Updatable); ok {
		v.Update()
	}

	defer makeWritable(r.buffer)()
	defer restoreCurorPosition(r.buffer.api.CurrentWindow())()

	r.buffer.SetLines(r.view.Lines())
	r.buffer.Freeze()
}

type Renderer struct {
}

func (r *Renderer) Attach(b *Buffer, view View) {
	vr := ViewRenderer{buffer: b, view: view}

	b.Freeze()

	bo := b.Options
	bo.SetHidden(BufferHiddenHide)
	bo.SetType(BufferTypeNoFile)
	bo.SetListed(false)

	view.Attach(vr)

	b.Options.SetFileType(view.FileType())
	if v, ok := view.(Initializable); ok {
		v.Initialize(b, b.api)
	}

	if v, ok := view.(disposables.Disposable); ok {
		disposed := false

		b.On(EventBufWipeout, func() {
			disposed = true
			v.Dispose()
		})

		// TODO debug why this fallback is required
		b.On(EventBufLeave, func() {
			go func() {
				time.Sleep(200 * time.Millisecond)
				if !disposed && !b.Exists() {
					disposed = true
					v.Dispose()
				}
			}()
		})

	}

	vr.ShouldRender()
}

func makeWritable(b *Buffer) func() {
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

func restoreCurorPosition(win *Window) func() {
	cur := win.Cursor()

	return func() {
		win.SetCursor(Cursor{cur[0], 0})
	}
}
