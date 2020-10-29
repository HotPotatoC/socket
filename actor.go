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

// Send write message to this actor
func (actor *Actor) Send(message []byte) error {
	return nil
}

// Close this function probably will working well
// but sometime it's not safe e.g: the connection already
// closed. Will return bool if connection closing correctly
func (actor *Actor) Close(message []byte) (bool, error) {
	return false, nil
}
