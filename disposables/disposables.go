package disposables

import "log"

type Disposable interface {
	Dispose()
}

var _ Disposable = (*Func)(nil)

type Func struct {
	dispose func()
}

func (h *Func) Dispose() {
	h.dispose()
}

func New(fn func()) Disposable {
	return &Func{
		dispose: fn,
	}
}

var _ Disposable = (*Collection)(nil)

type Collection struct {
	content []Disposable
}

func NewCollection() *Collection {
	return &Collection{
		content: []Disposable{},
	}
}

func (ds *Collection) Add(d Disposable) {
	ds.content = append(ds.content, d)
}

func (ds *Collection) Dispose() {
	for _, d := range ds.content {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("dispose error: %v", err)
				}
			}()
			d.Dispose()
		}()
	}

	ds.content = []Disposable{}
}

