package socket_test

import (
	"testing"

	"github.com/Stalync/socket"
)

func TestCreateWebSocketConnection(t *testing.T) {
	socket := socket.CreateWebSocket()
	socket.Listen(8080)
}

func TestCreateWebSocketConnectionWithLogger(t *testing.T) {
	socket := socket.CreateWebSocket()
	socket.SetLogger()
	socket.Listen(8080)
}

func TestConnectedEventWithGreeting(t *testing.T) {
	socket := socket.CreateWebSocket()

	socket.Callback(func(c *socket.Context) error {
		var err error
		if c.Event().Type() != "connected" {
			return err
		}
		sender := c.Sender()
		// You can do this
		err = sender.Send("Welcome User")

		// or like this
		err = socket.SendTo(sender.ID(), "Welcome User")
	})

	socket.Listen(8080)
}

func TestSendCloseAndForceCloseConnection(t *testing.T) {
	socket := socket.CreateWebSocket()

	socket.Callback(func(c *socket.Context) error {
		var err error
		if c.Event().Type() != "connected" {
			return err
		}
		sender := c.Sender()
		sender.Send("Your connection is closed")
		ok, err := sender.ForceClose()
	})

	socket.Listen(8080)
}

func TestAnyEventWithParseMessage(t *testing.T) {
	socket := socket.CreateWebSocket()
	defer socket.Listen(8080)

	socket.Callback(func(c *socket.Context) error {
		var err error = nil
		var eType = c.Event().Type()

		if eType != socket.TypeConnected || eType != socket.TypeClose {
			return err
		}
		sender := c.Sender()
		messages := c.Message()
		if messages.Type == socket.TypeText {

			text := messages.Bytes()

			// Send To All
			err := socket.Send(text)

			// Send To All Without me
			socket.SendWithFilter(func(s *socket.Actor, callback func(string) error) {
				s.Attr.(Attr)
				if s.ID != sender.ID() {
					err = callback(s.ID())
				}
			}, messages.Bytes(), []byte(`bla blac`), []byte(`bla blac 2`))

			// Or with async send
			socket.SendAsyncWithFilter(func(s *socket.Actor, callback func(string) error) {
				s.Attr.(Attr)
				if s.ID != sender.ID() {
					err = callback(s.ID())
				}
			}, messages.Bytes(), []byte(`bla blac`), []byte(`bla blac 2`))
		}
	})
}

type Attr struct {
	Room string
}
