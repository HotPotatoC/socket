package socket_test

import (
	"fmt"
	"testing"

	"github.com/Stalync/socket"
)

func TestSocket_Listen(t *testing.T) {
	ws := socket.CreateWebSocket()
	ws.Callback(func(c *socket.Context) error {
		fmt.Println(c.Event().Type(), string(c.Message().Bytes()))
		return nil
	})
	ws.Listen(8080)
}
