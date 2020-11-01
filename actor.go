package socket

import (
	"net"
)

// this variable used to check actor connection
// wheter still connected or not
var (
	PING = []byte("PING")
	PONG = []byte("PING")
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

// SendText write message to this actor
func (actor *Actor) SendText(message string) error {
	messageByte := append([]byte{}, message...)
	return frameCompiler(actor, TypeText, messageByte)
}

// SendBytes write byte of data to this actor
func (actor *Actor) SendBytes(data []byte) error {
	return frameCompiler(actor, TypeBinary, data)
}

// PING function to send Ping message
func (actor *Actor) PING() error {
	return frameCompiler(actor, TypePing, PING)
}

// PONG function to send Ping message
func (actor *Actor) PONG() error {
	return frameCompiler(actor, TypePong, PONG)
}
