package socket

import "github.com/gobwas/ws"

// Close no meaning close function
func (s *Socket) Close(a *Actor) (err error) {
	s.actors.Delete(a.ID())
	err = s.CloseNoMeaningYet(a, "")
	if err == nil {
		err = (*a.conn).Close()
	}
	return
}

// CloseNormalClosure is used to close actor connection
// with defined status
func (s *Socket) CloseNormalClosure(a *Actor, message string) (err error) {
	s.actors.Delete(a.ID())
	reason := ws.NewCloseFrameBody(ws.StatusNormalClosure, message)
	frame := ws.NewCloseFrame(reason)
	err = ws.WriteFrame((*a.conn), frame)

	if err == nil {
		err = (*a.conn).Close()
	}
	return
}

// CloseGoingAway is used to close actor connection
// with defined status
func (s *Socket) CloseGoingAway(a *Actor, message string) (err error) {
	s.actors.Delete(a.ID())
	reason := ws.NewCloseFrameBody(ws.StatusGoingAway, message)
	frame := ws.NewCloseFrame(reason)
	err = ws.WriteFrame((*a.conn), frame)

	if err == nil {
		err = (*a.conn).Close()
	}
	return
}

// CloseProtocolError is used to close actor connection
// with defined status
func (s *Socket) CloseProtocolError(a *Actor, message string) (err error) {
	s.actors.Delete(a.ID())
	reason := ws.NewCloseFrameBody(ws.StatusProtocolError, message)
	frame := ws.NewCloseFrame(reason)
	err = ws.WriteFrame((*a.conn), frame)

	if err == nil {
		err = (*a.conn).Close()
	}
	return
}

// CloseUnsupportedData is used to close actor connection
// with defined status
func (s *Socket) CloseUnsupportedData(a *Actor, message string) (err error) {
	s.actors.Delete(a.ID())
	reason := ws.NewCloseFrameBody(ws.StatusUnsupportedData, message)
	frame := ws.NewCloseFrame(reason)
	err = ws.WriteFrame((*a.conn), frame)

	if err == nil {
		err = (*a.conn).Close()
	}
	return
}

// CloseNoMeaningYet is used to close actor connection
// with defined status
func (s *Socket) CloseNoMeaningYet(a *Actor, message string) (err error) {
	s.actors.Delete(a.ID())
	reason := ws.NewCloseFrameBody(ws.StatusNoMeaningYet, message)
	frame := ws.NewCloseFrame(reason)
	err = ws.WriteFrame((*a.conn), frame)

	if err == nil {
		err = (*a.conn).Close()
	}
	return
}

// CloseInvalidFramePayloadData is used to close actor connection
// with defined status
func (s *Socket) CloseInvalidFramePayloadData(a *Actor, message string) (err error) {
	s.actors.Delete(a.ID())
	reason := ws.NewCloseFrameBody(ws.StatusInvalidFramePayloadData, message)
	frame := ws.NewCloseFrame(reason)
	err = ws.WriteFrame((*a.conn), frame)

	if err == nil {
		err = (*a.conn).Close()
	}
	return
}

// ClosePolicyViolation is used to close actor connection
// with defined status
func (s *Socket) ClosePolicyViolation(a *Actor, message string) (err error) {
	s.actors.Delete(a.ID())
	reason := ws.NewCloseFrameBody(ws.StatusPolicyViolation, message)
	frame := ws.NewCloseFrame(reason)
	err = ws.WriteFrame((*a.conn), frame)

	if err == nil {
		err = (*a.conn).Close()
	}
	return
}

// CloseMessageTooBig is used to close actor connection
// with defined status
func (s *Socket) CloseMessageTooBig(a *Actor, message string) (err error) {
	s.actors.Delete(a.ID())
	reason := ws.NewCloseFrameBody(ws.StatusMessageTooBig, message)
	frame := ws.NewCloseFrame(reason)
	err = ws.WriteFrame((*a.conn), frame)

	if err == nil {
		err = (*a.conn).Close()
	}
	return
}

// CloseMandatoryExt is used to close actor connection
// with defined status
func (s *Socket) CloseMandatoryExt(a *Actor, message string) (err error) {
	s.actors.Delete(a.ID())
	reason := ws.NewCloseFrameBody(ws.StatusMandatoryExt, message)
	frame := ws.NewCloseFrame(reason)
	err = ws.WriteFrame((*a.conn), frame)

	if err == nil {
		err = (*a.conn).Close()
	}
	return
}

// CloseInternalServerError is used to close actor connection
// with defined status
func (s *Socket) CloseInternalServerError(a *Actor, message string) (err error) {
	s.actors.Delete(a.ID())
	reason := ws.NewCloseFrameBody(ws.StatusInternalServerError, message)
	frame := ws.NewCloseFrame(reason)
	err = ws.WriteFrame((*a.conn), frame)

	if err == nil {
		err = (*a.conn).Close()
	}
	return
}

// CloseTLSHandshake is used to close actor connection
// with defined status
func (s *Socket) CloseTLSHandshake(a *Actor, message string) (err error) {
	s.actors.Delete(a.ID())
	reason := ws.NewCloseFrameBody(ws.StatusTLSHandshake, message)
	frame := ws.NewCloseFrame(reason)
	err = ws.WriteFrame((*a.conn), frame)

	if err == nil {
		err = (*a.conn).Close()
	}
	return
}

// Callback used to set handler of incoming message
func (s *Socket) Callback(cb func(c *Context) error) {
	s.cb = func(c *Context) error {
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

// SendByteTo function that can enable to send message
// to other connected client
func (s *Socket) SendByteTo(id string, data []byte) error {
	var actor *Actor
	var found bool
	if actor, found = s.actors.Read(id); !found {
		return errIDNotFound
	}
	return actor.SendBytes(data)
}
