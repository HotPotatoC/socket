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
	if port <= 1024 {
		return errors.New("Port number should be larger than 1024 and not being used")
	}

	addr := fmt.Sprintf("localhost:%v", port)
	ln, err := net.Listen("tcp", addr)

	if err != nil {
		return err
	}
	u := createUpgrader(s.config)
	s.config.HostWhitelist = append(s.config.HostWhitelist, addr)

	for {
		conn, err := ln.Accept()
		if err == nil {
			_, err = u.Upgrade(conn)
			if err == nil {
				currentActor := s.registerActor(conn)
				s.serveActorMessage(currentActor)
			}
		}

		if err != nil {
			conn.Close()
		}
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

	}()
}

// Callback used to set handler of incoming message
func (s *Socket) Callback(cb func(c *Context) error) {
	s.cb = cb
}

func createUpgrader(config *Config) *ws.Upgrader {
	return &ws.Upgrader{
		OnHost: func(host []byte) error {
			hostString := string(host)
			for _, y := range config.HostWhitelist {
				if y == hostString {
					return nil
				}
			}
			return errors.New("Host not allowed")
		},
	}
}
