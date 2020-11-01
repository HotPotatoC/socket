package workers

import (
	"github.com/Stalync/socket"
	"github.com/Stalync/socket/workers/wpond"
)

// Worker this interface defined Worker behaviour
// so, we can use many backend workers API
type Worker interface {

	// Callback can only be called when initializing the worker
	// and should not be nil
	// this function will responsible to what you want to do
	// with passed data from worker pool
	Callback(func(interface{}))

	// CallbackOnError same as Callback but can set to nil
	// this function set what you want to do when data/task
	// is not submited or not running properly
	CallbackOnError(func(interface{}))

	// Submit this function supposed to handle submit data
	// parameter can be list
	Submit(...interface{})
}

func CreateWorkerPond(c *socket.Config) Worker {
	worker := wpond.WorkerPond{}
	worker.Init(c.WorkerPool, c.BufferedTask)
	return &worker
}
