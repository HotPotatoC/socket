package socket

import "time"

// Config storing config for incoming message callback
type Config struct {
	Timeout       time.Duration
	HostWhitelist []string
}

var DefaultConfig = Config{
	Timeout:       time.Second * 4,
	HostWhitelist: make([]string, 0),
}
