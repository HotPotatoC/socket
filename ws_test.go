package socket_test

// func TestCreateWebSocketConnection(t *testing.T) {
// 	ws := socket.CreateWebSocket()
// 	ws.Listen(8080)
// }

// func TestCreateWebSocketConnectionWithLogger(t *testing.T) {
// 	ws := socket.CreateWebSocket()
// 	ws.SetLogger()
// 	ws.Listen(8080)
// }

// func TestConnectedEventWithGreeting(t *testing.T) {
// 	ws := socket.CreateWebSocket()

// 	ws.Callback(func(c *socket.Context) error {
// 		var err error
// 		if c.Event().Type() != "connected" {
// 			return
// 		}
// 		sender := c.Sender()
// 		// You can do this
// 		err = sender.Send("Welcome User")

// 		// or like this
// 		err = ws.SendTo(sender.ID(), "Welcome User")
// 	})

// 	ws.Listen(8080)
// }

// func TestSendCloseAndCloseConnection(t *testing.T) {
// 	ws := socket.CreateWebSocket()

// 	ws.Callback(func(c *socket.Context) error {
// 		var err error
// 		if c.Event().Type() != "connected" {
// 			return
// 		}
// 		sender := c.Sender()
// 		sender.Send("Your connection is closed")
// 		err := sender.CloseWithMessage(ws.StatusNormalClosure, "You are doing it well")
// 	})

// 	ws.Listen(8080)
// }

// func TestAnyEventWithParseMessage(t *testing.T) {
// 	ws := socket.CreateWebSocket()
// 	defer ws.Listen(8080)

// 	ws.Callback(func(c *socket.Context) error {
// 		var err error = nil
// 		var eType = c.Event().Type()

// 		if eType != socket.TypeConnected || eType != socket.TypeClose {
// 			return err
// 		}
// 		sender := c.Sender()
// 		messages := c.Message()
// 		if messages.Type == socket.TypeText {

// 			text := messages.Bytes()

// 			// Send To All
// 			err := ws.Send(text)

// 			// Send To All Without me
// 			ws.SendWithFilter(func(s *socket.Actor, callback func(string) error) {
// 				s.Attr.(Attr)
// 				if s.ID != sender.ID() {
// 					err = callback(s.ID())
// 				}
// 			}, messages.Bytes(), []byte(`bla blac`), []byte(`bla blac 2`))

// 			// Or with async send
// 			ws.SendAsyncWithFilter(func(s *socket.Actor, callback func(string) error) {
// 				s.Attr.(Attr)
// 				if s.ID != sender.ID() {
// 					err = callback(s.ID())
// 				}
// 			}, messages.Bytes(), []byte(`bla blac`), []byte(`bla blac 2`))
// 		}
// 	})
// }

// type Attr struct {
// 	Room string
// }
