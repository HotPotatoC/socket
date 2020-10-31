package socket_test

import (
	"fmt"
	"testing"

	"github.com/Stalync/socket"
)

func TestSocket_Listen(t *testing.T) {
	server := socket.CreateWebSocket()

	server.Callback(func(c *socket.Context) (err error) {

		if ok, _ := c.Event().Type().Eq(socket.TypeText); ok {
			return c.Sender().SendText(c.Message().String())
		}

		return nil
	})

	fmt.Println(server.Listen(8000))
}
