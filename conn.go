package socket

import (
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/gobwas/ws"
)

// Socket struct for storing connection information
type Socket struct {
	port int

	// list of connected client
	actors map[string]*Actor

	config *Config
}

// CreateWebSocket constructor to create socket connection
func CreateWebSocket() *Socket {
	return &Socket{
		port:   8000,
		actors: make(map[string]*Actor),
	}
}

// Listen function that loop every time to catch new connection
func (s *Socket) Listen(port int) error {
	if port <= 1024 {
		return errors.New("Port number should be larger than 1024 and not being used")
	}
	ln, err := createNetListener(port)
	if err != nil {
		return err
	}
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

func createNetListener(port int) (*net.Listener, error) {
	ps := strconv.Itoa(port)
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", ps))
	return &ln, err
}
