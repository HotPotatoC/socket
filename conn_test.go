package socket_test

import (
	"fmt"
	"testing"

	"github.com/Stalync/socket"
)

func TestSocket_Listen(t *testing.T) {
	ws := socket.CreateWebSocket()
	ws.Callback(func(c *socket.Context) error {
		if ok, err := c.Event().Type().Eq(socket.TypeConnected); err == nil && ok {
			fmt.Printf("Someone connected with id: '%v'", c.Sender().ID())
		}
		return nil
	})
	ws.Listen(8080)
}
