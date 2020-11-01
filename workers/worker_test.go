package workers_test

import (
	"fmt"
	"testing"

	"github.com/Stalync/socket"
	"github.com/Stalync/socket/workers"
)

func TestMain(t *testing.T) {
	config := socket.DefaultConfig
	pool := workers.CreateWorkerPond(&config)
	go func() {
		for i := 0; i < 10000; i++ {
			pool.Submit(i)
		}
	}()
	pool.Callback(func(param interface{}) {
		fmt.Println(param.(int))
	})
	pool.CallbackOnError(func(param interface{}) {
		fmt.Println("err", param)
	})

	// Loop forever, but you can use counter to check whether
	// the all submited task is done
	for {

	}
}
