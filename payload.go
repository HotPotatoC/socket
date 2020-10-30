package socket

import (
	"context"

	"github.com/gobwas/ws"
)

// Message stored data of incoming signal from client
type Message struct {
	data *[]byte
}

// Length return length of incoming data
func (msg *Message) Length() int {
	return len(*msg.data)
}

// Bytes return array of bytes incoming data
func (msg *Message) Bytes() []byte {
	return *msg.data
}

// Event describe of incoming event type
type Event struct {
	header *ws.Header
}

// Type return event type
func (e *Event) Type() TypeCode {
	return TypeCode(e.header.OpCode)
}

// Header return event header
func (e *Event) Header() ws.Header {
	return *e.header
}

// Context stored info from client
type Context struct {
	message Message
	sender  Actor
	event   Event
	timeout context.Context
}

func createContext(config *Config) *Context {
	ctx, _ := context.WithTimeout(
		context.Background(),
		config.Timeout,
	)
	return &Context{
		timeout: ctx,
		event: Event{
			&ws.Header{},
		},
		message: Message{
			data: &[]byte{},
		},
	}
}

// Event return event context
func (c *Context) Event() *Event {
	return &c.event
}

// Message return message context
func (c *Context) Message() *Message {
	return &c.message
}

// Sender return Actor from this context
func (c *Context) Sender() *Actor {
	return &c.sender
}

// Ctx return parrent Context
func (c *Context) Ctx() context.Context {
	return c.timeout
}
