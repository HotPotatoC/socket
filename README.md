# Stalync Socket
> High-performance Standalone WebSocket, built for powering Stalync WS cluster based on [ws](https://github.com/gobwas/ws)

This project not done yet, and we still working to improving this. If you have any thought how to improve this, feel free to create pull request.

# Features

- [X] Simple API (but still low level)
- [X] Send message to other client (Sync)
- [X] Close connection
- [ ] Graceful shutdown
- [ ] Data encrypt
- [ ] Data response information (Ack, etc)
- [ ] Send with worker pool (Async)
- [ ] Non-blocking IO (Adapter for various TCP Framework)
- [ ] Rate limiting
- [ ] Container instant WebSocket
- [ ] Resource pooling
- [ ] Clustering/Scaling
- [ ] Authentication
  - JWT Token
  - Host Whitelist
  - Cookie parser
- [ ] Alerting
  - Logger

# Usage
The example below still terrible but you can customize by yourself
```go
package main

import (
	"fmt"
	"testing"

	"github.com/Stalync/socket"
	"github.com/gobwas/ws"
)

func main() {
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

			if message == "broadcast" {

				server.IterateEachActor(func(actor *socket.Actor) {

					if actor.ID() == c.Sender().ID() {
						return
					}

					err := actor.SendText(fmt.Sprintf("Hellow from %v", c.Sender().ID()))

					if err != nil {
						fmt.Println(fmt.Sprintf("Send from %v to %v failed with message: %v",
							c.Sender().ID(),
							actor.ID(),
							err,
						))
					}

				})
				return
			}
			err = c.Sender().SendText(message)
			break

		case socket.TypeDisconnected:
			fmt.Println(fmt.Sprintf("Master %v is disconnected", c.Sender().ID()))
			break
		}

		return
	})

	fmt.Println(server.Listen(8080))
}
```
using [wscat](https://github.com/websockets/wscat) to test connection

```bash
wscat -c ws://localhost:8000
```

# Need to know

### List of known access event type
- TypeContinuation
- TypeConnected
- TypeDisconnected
- TypePing
- TypePong  

### List of known data type
- TypeText
- TypeBinary
