# Stalync Socket
> High-performance Standalone WebSocket, built for powering Stalync WS cluster based on [ws](https://github.com/gobwas/ws)

This project not done yet, and we still working to improving this. If you have any thought how to improve this, feel free to create pull request.

# Features

- [X] Simple API (but still low level)
- [X] Send message to other client (Sync)
- [X] Close connection
- [ ] Resource pooling
- [ ] Clustering/Scaling
- [ ] Authentication
  - JWT Token
  - Host Whitelist
  - Cookie parser
- [ ] Send with worker pool (Async)

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
    
    eType := c.Event().Type()
		if ok, _ := eType.Eq(socket.TypeConnected); ok {

			fmt.Println("Someone connected with id: " + c.Sender().ID())
			return c.Sender().SendText("Welcome user")

		}
		if ok, _ := eType.Eq(socket.TypeDisconnected); c.Message().String() == "exit" || ok {

			// Close function should be called, it will handle delete session in internal server
			err = server.CloseByActorWithMessage(c.Sender(), ws.StatusNormalClosure, "Byee Human")
			return nil
		}

		return nil
	})

	fmt.Println(server.Listen(8000))
}
```
