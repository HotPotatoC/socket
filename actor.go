package socket

import (
	"net"
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
	return frameBuilderAndSender(actor, TypeText, messageByte)
}
