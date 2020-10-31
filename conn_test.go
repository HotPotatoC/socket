package socket_test

import (
	"fmt"
	"testing"

	"github.com/Stalync/socket"
)

func TestSocket_Listen(t *testing.T) {
	server := socket.CreateWebSocket()

	server.Callback(func(c *socket.Context) (err error) {

		switch c.Event().Type() {
		case socket.TypeConnected:
			fmt.Println(fmt.Sprintf("Master %v is connected", c.Sender().ID()))
			err = c.Sender().SendText("Hellow Master")
			break
		case socket.TypeText:
			message := c.Message().String()
			if message == "exit" {
				return server.CloseNormalClosure(c.Sender(), "Byeeee Master")
			}
			err = c.Sender().SendText(message)
			break
		case socket.TypeDisconnected:
			fmt.Println(fmt.Sprintf("Master %v is disconnected", c.Sender().ID()))
			break
		}

		return
	})

	fmt.Println(server.Listen(8000))
}
