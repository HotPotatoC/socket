package socket

import "net"

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
