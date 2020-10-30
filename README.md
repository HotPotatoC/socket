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

	server.Callback(func(c *socket.Context) error {
    
    	eType := c.Event().Type()
	if ok, _ := eType.Eq(socket.TypeConnected); ok {

		fmt.Println("Someone connected with id: " + c.Sender().ID())
		return c.Sender().SendText("Welcome user")

	}

	if ok, _ := eType.Eq(socket.TypeDisconnected); c.Message().String() == "exit" || ok {

		// Close function should be called, it will handle delete session in internal server
		err = server.CloseByActorWithMessage(c.Sender(), ws.StatusNormalClosure, "Byee Human")
		return err
	}
	return c.Sender().SendText(c.Message().String())
	})

	fmt.Println(server.Listen(8000))
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

### List of known access event type
- TypeText
- TypeBinary
