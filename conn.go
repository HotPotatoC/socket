package socket

import (
	"errors"
	"fmt"
	"net"

	"github.com/gobwas/ws"
)

// Socket struct for storing connection information
type Socket struct {
	port int

	// list of connected client
	actors map[string]*Actor

	config *Config

	cb func(c *Context) error
}

// CreateWebSocket constructor to create socket connection
func CreateWebSocket() *Socket {
	return &Socket{
		port:   8000,
		actors: make(map[string]*Actor),
		config: &DefaultConfig,
	}
}

// Listen function that loop every time to catch new connection
func (s *Socket) Listen(port int) error {
	if s.cb == nil {
		return errors.New("Callback should not be nill")
	}
	if port <= 1024 {
		return errors.New("Port number should be larger than 1024 and not being used")
	}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))

	if err != nil {
		return err
	}
	u := createUpgrader(s.config)

	for {
		conn, err := ln.Accept()
		var currentActor *Actor

		if err == nil {
			_, err = u.Upgrade(conn)
			if err == nil {
				currentActor = s.registerActor(conn)
				s.cb(connectedContext(currentActor))
			}
		}
		if err != nil {
			conn.Close()
			continue
		}
		s.serveActorMessage(currentActor)
	}
}

func (s *Socket) registerActor(conn net.Conn) *Actor {
	id := generateKey(s.config.UIDLength)
	_, found := s.actors[id]
	for found {
		id = generateKey(s.config.UIDLength)
		_, found = s.actors[id]
	}
	s.actors[id] = &Actor{
		id:   &id,
		conn: &conn,
	}
	return s.actors[id]
}

// serveActorMessage function that listen incoming message
// of connected actors. and call callback when message accepted
func (s *Socket) serveActorMessage(a *Actor) {
	go func() {
		for {
			ctx, err := s.contextBuilder(a)
			if err == nil {
				s.cb(ctx)
			}
		}
	}()
}

// CloseByActorWithMessage this function supposed
// to close connection with status code and message
func (s *Socket) CloseByActorWithMessage(a *Actor, code ws.StatusCode, message string) (err error) {
	delete(s.actors, a.ID())
	err = frameBuilderAndSender(a, TypeDisconnected, []byte(message), code)
	if err == nil {
		err = (*a.conn).Close()
	}
	return
}

// CloseByIDWithMessage this function supposed
// to close connection with status code and message
func (s *Socket) CloseByIDWithMessage(id string, code ws.StatusCode, message string) (err error) {
	actor, found := s.actors[id]
	delete(s.actors, id)

	if found {
		err = frameBuilderAndSender(actor, TypeDisconnected, []byte(message), code)
		if err == nil {
			err = (*actor.conn).Close()
		}
	}
	return
}

// Callback used to set handler of incoming message
func (s *Socket) Callback(cb func(c *Context) error) {
	s.cb = cb
}

func createUpgrader(config *Config) *ws.Upgrader {
	return &ws.Upgrader{}
}

func handleWhitelistHost(config *Config, host []byte) error {
	hostString := string(host)
	for _, y := range config.hostWhitelist {
		if y == hostString {
			return nil
		}
	}
	return errors.New("Host not allowed")
}
