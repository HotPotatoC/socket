package socket

import "time"

// Config storing config for incoming message callback
type Config struct {
	Timeout       time.Duration
	HostWhitelist []string
	UIDLength     int
}

// DefaultConfig setting up for default connection config
var DefaultConfig = Config{
	Timeout:       time.Second * 4,
	HostWhitelist: make([]string, 0),
	UIDLength:     16,
}
