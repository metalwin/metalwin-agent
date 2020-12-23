package client

import (
	"fmt"

	libvirt "github.com/metalwin/metalwin-agent/internal"
	"github.com/metalwin/metalwin-agent/pkg/config"
)

// Client to make requests to the libvirt daemon.
type Client struct {
	conn   libvirt.LibvirtClient
	config config.ConnectionConfig
	Domain *domainService
}

// NewClient generates a client to interact with the libvirt daemon in
// the host machine.
func NewClient(config config.ConnectionConfig) (*Client, error) {
	conn, err := libvirt.NewConnection(config.String())
	if err != nil {
		return nil, err
	}
	domainService := domainService{conn: conn}
	return &Client{conn, config, &domainService}, nil
}

// HypervisorVersion returns the version of the hypervisor
func (c *Client) HypervisorVersion() (string, error) {
	v, err := c.conn.GetVersion()
	if err != nil {
		return "", err
	}
	return transformVersion(v), nil
}

// LibvirtDaemonVersion returns the version of the libvirt daemon where the host is running
func (c *Client) LibvirtDaemonVersion() (string, error) {
	v, err := c.conn.GetLibVersion()
	if err != nil {
		return "", err
	}
	return transformVersion(v), nil
}

func transformVersion(version uint32) string {
	major := version / 1_000_000
	version -= major * 1_000_000
	minor := version / 1_000
	version -= minor * 1_000
	return fmt.Sprintf("%d.%d.%d", major, minor, version)
}

// Close releases all resources from this client
func (c *Client) Close() (int, error) {
	return c.conn.Close()
}
