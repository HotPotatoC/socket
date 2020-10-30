package socket

import "time"

// Config storing config for incoming message callback
type Config struct {
	Timeout       time.Duration
	hostWhitelist []string
	UIDLength     int
}

func (c *Config) pushHostWhitelist(h string) {
	c.hostWhitelist = append(c.hostWhitelist, h)
}

// DefaultConfig setting up for default connection config
var DefaultConfig = Config{
	Timeout:       time.Second * 4,
	hostWhitelist: make([]string, 0),
	UIDLength:     16,
}
