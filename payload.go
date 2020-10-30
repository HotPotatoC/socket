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
	h    *ws.Header
	code *TypeCode
}

// Type return event type
func (e *Event) Type() *TypeCode {
	return e.code
}

func (e *Event) header() *ws.Header {
	return e.h
}

// Context stored info from client
type Context struct {
	message *Message
	sender  *Actor
	event   *Event
	timeout context.Context
}

func createContext(config *Config) *Context {
	ctx, _ := context.WithTimeout(
		context.Background(),
		config.Timeout,
	)
	return &Context{
		timeout: ctx,
		event: &Event{
			h: &ws.Header{},
		},
		message: &Message{
			data: &[]byte{},
		},
	}
}

// Event return event context
func (c *Context) Event() *Event {
	return c.event
}

// Message return message context
func (c *Context) Message() *Message {
	return c.message
}

// Sender return Actor from this context
func (c *Context) Sender() *Actor {
	return c.sender
}

// Ctx return parrent Context
func (c *Context) Ctx() context.Context {
	return c.timeout
}

var connectedContext = func() *Context {
	c := createContext(&DefaultConfig)
	connected := TypeConnected
	c.event.code = &connected
	return c
}()
