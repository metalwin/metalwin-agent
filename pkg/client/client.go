package client

import (
	libvirt "github.com/metalwin/metalwin-agent/internal"
	"github.com/metalwin/metalwin-agent/pkg/config"
)

// Client to make requests to the libvirt daemon.
type Client struct {
	conn   libvirt.LibvirtClient
	config config.ConnectionConfig
}

// NewClient generates a client to interact with the libvirt daemon in
// the host machine.
func NewClient(config config.ConnectionConfig) (*Client, error) {
	conn, err := libvirt.NewConnection(config.String())
	if err != nil {
		return nil, err
	}
	return &Client{conn, config}, nil
}

// Version of the hypervisor
func (c *Client) Version() (uint32, error) {
	return c.conn.GetVersion()
}
