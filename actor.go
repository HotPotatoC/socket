package socket

import (
	"net"

	"github.com/gobwas/ws"
)

// Actor struct containing information of actor
// Actor can be said as an connected client
type Actor struct {
	id   *string
	conn *net.Conn
}

// ID return id information
func (actor *Actor) ID() string {
	return *actor.id
}

// Send write message to this actor
func (actor *Actor) Send(message []byte) error {
	return nil
}

// CloseWithMessage this function supposed to close connection
// with status code and message
func (actor *Actor) CloseWithMessage(code ws.StatusCode, message string) error {
	return nil
}
