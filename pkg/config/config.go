package config

import "fmt"

// ConnectionConfig used to make a connection to the host libvirt daemon
type ConnectionConfig struct {
	Driver string
	Path   string
}

// String transform this confuguration to a readable string
func (c *ConnectionConfig) String() string {
	return fmt.Sprintf("%s:///%s", c.Driver, c.Path)
}
