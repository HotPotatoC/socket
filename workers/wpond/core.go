package wpond

import (
	"github.com/alitto/pond"
)

// WorkerPond Worker wrapper for Pond API
type WorkerPond struct {
	pool      *pond.WorkerPool
	cb        func(interface{})
	cbOnError func(interface{})
}

// Init function to init pond API
func (w *WorkerPond) Init(worker, bufferedTask int, options ...pond.Option) {
	w.pool = pond.New(worker, bufferedTask, options...)
}

// Callback set function to process data
func (w *WorkerPond) Callback(cb func(interface{})) {
	w.cb = cb
}

// CallbackOnError set function to process error data
func (w *WorkerPond) CallbackOnError(cb func(interface{})) {
	w.cbOnError = cb
}

// Submit function to submit parameter
func (w *WorkerPond) Submit(params ...interface{}) {
	if w.cb != nil && w.cbOnError != nil {
		for _, y := range params {
			func(param interface{}) {
				submited := w.pool.TrySubmit(func() {
					w.cb(param)
				})
				if !submited {
					w.cbOnError(param)
				}
			}(y)
		}
	} else {
		panic("Callback and CallbackOnError should not be nil")
	}
}
