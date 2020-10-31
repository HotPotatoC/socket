package socket

import (
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/gobwas/ws"
)

var (
	errCbNill           = errors.New("Callback should not be nill")
	errPortNotValid     = errors.New("Port number should be larger than 1024 and not being used")
	errNotAllowed       = errors.New("Host not allowed")
	errIDNotFound       = errors.New("ID Not found")
	errTypeNotSupported = errors.New("Type not supported")
)

// ActorMap this map only storing actor and actor id
// but with more safety
type ActorMap struct {
	sync.RWMutex
	maps map[string]*Actor
}

// Insert actor to map
func (m *ActorMap) Insert(id string, a *Actor) {
	m.Lock()
	m.maps[id] = a
	m.Unlock()
}

// Delete actor from map
func (m *ActorMap) Delete(id string) {
	m.Lock()
	delete(m.maps, id)
	m.Unlock()
}

// Read function that read actor from map
func (m *ActorMap) Read(id string) (*Actor, bool) {
	m.Lock()
	actor, found := m.maps[id]
	m.Unlock()
	return actor, found
}

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
		return errCbNill
	}
	if port <= 1024 {
		return errPortNotValid
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
	return errNotAllowed
}

// CloseByActorWithMessage this function supposed
// to close connection with status code and message
func (s *Socket) CloseByActorWithMessage(a *Actor, code ws.StatusCode, message string) (err error) {
	delete(s.actors, a.ID())
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
	if a, found := s.actors[id]; found {
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
	if actor, found = s.actors[id]; !found {
		return errIDNotFound
	}
	return actor.SendText(message)
}
