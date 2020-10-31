package socket

import "github.com/gobwas/ws"

// CloseByActorWithMessage this function supposed
// to close connection with status code and message
func (s *Socket) CloseByActorWithMessage(a *Actor, code ws.StatusCode, message string) (err error) {
	s.actors.Delete(a.ID())
	messageByte := append([]byte{}, message...)
	err = frameBuilderAndSender(a, TypeDisconnected, messageByte, code)
	if err == nil {
		err = (*a.conn).Close()
	}
	return
}

// CloseByIDWithMessage this function supposed
// to close connection with status code and message
func (s *Socket) CloseByIDWithMessage(id string, code ws.StatusCode, message string) error {
	if a, found := s.actors.Read(id); found {
		return s.CloseByActorWithMessage(a, code, message)
	}
	return errIDNotFound
}

// Callback used to set handler of incoming message
func (s *Socket) Callback(cb func(c *Context) error) {
	s.cb = func(c *Context) error {
		defer func() {
			if *c.event.code == TypeDisconnected {
				s.CloseByIDWithMessage(*c.sender.id, ws.StatusNoMeaningYet, "")
			}
		}()
		return cb(c)
	}
}

// SendTextTo function that can enable to send message
// to other connected client
func (s *Socket) SendTextTo(id, message string) error {
	var actor *Actor
	var found bool
	if actor, found = s.actors.Read(id); !found {
		return errIDNotFound
	}
	return actor.SendText(message)
}
