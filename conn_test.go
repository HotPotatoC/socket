package socket_test

import (
	"fmt"
	"testing"

	"github.com/Stalync/socket"
	"github.com/gobwas/ws"
)

func TestSocket_Listen(t *testing.T) {
	server := socket.CreateWebSocket()

	server.Callback(func(c *socket.Context) (err error) {

		if ok, _ := c.Event().Type().Eq(socket.TypeConnected); ok {

			fmt.Println("Someone connected with id: " + c.Sender().ID())
			return c.Sender().SendText("Welcome user")

		}
		if ok, _ := c.Event().Type().Eq(socket.TypeDisconnected); c.Message().String() == "exit" || ok {

			// Close function should be called, it will handle delete session in internal server
			err = server.CloseByActorWithMessage(c.Sender(), ws.StatusNormalClosure, "Byee Human")
			return nil
		}

		return nil
	})

	fmt.Println(server.Listen(8000))
}
