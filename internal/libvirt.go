package internal

import (
	"errors"

	libvirt "libvirt.org/libvirt-go"
)

var errAgentNotConnected = errors.New("agent is not connected to libvirt daemon")

// LibvirtConnection holds a connection to the libvirt daemon
type LibvirtConnection struct {
	conn *libvirt.Connect
	addr string
}

// LibvirtClient implements the operations required for this client to make.
type LibvirtClient interface {
	Close() (int, error)
	GetVersion() (uint32, error)
	GetLibVersion() (uint32, error)
	ListAllDomains() ([]Domain, error)
}

type Domain interface {
	GetName() (string, error)
}

// NewConnection creates a connection to the libvirt daemon using the given
// connection address.
func NewConnection(address string) (LibvirtClient, error) {
	conn, err := libvirt.NewConnect(address)
	if err != nil {
		return nil, err
	}
	return &LibvirtConnection{conn, address}, nil
}

// Close the libvirt daemon connection
func (c *LibvirtConnection) Close() (int, error) {
	if c.conn == nil {
		return 0, nil
	}
	return c.conn.Close()
}

// GetVersion of the running hypervisor
func (c *LibvirtConnection) GetVersion() (uint32, error) {
	if c.conn == nil {
		return 0, errAgentNotConnected
	}
	return c.conn.GetVersion()
}

// GetLibVersion return the version of the libvirt used by the daemon
// running in the host
func (c *LibvirtConnection) GetLibVersion() (uint32, error) {
	if c.conn == nil {
		return 0, errAgentNotConnected
	}
	return c.conn.GetLibVersion()
}

// ListAllDomains retrieve the list of domains based on the given filter
func (c *LibvirtConnection) ListAllDomains() ([]Domain, error) {
	flags := libvirt.CONNECT_LIST_DOMAINS_PERSISTENT | libvirt.CONNECT_LIST_DOMAINS_TRANSIENT
	domains, err := c.conn.ListAllDomains(flags)
	if err != nil {
		return nil, err
	}
	result := make([]Domain, len(domains))
	for idx, domain := range domains {
		result[idx] = &domain
	}
	return result, nil
}
