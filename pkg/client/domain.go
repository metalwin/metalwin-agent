package client

import libvirt "github.com/metalwin/metalwin-agent/internal"

type domainService struct {
	conn libvirt.LibvirtClient
}

func (d *domainService) ListAll() ([]libvirt.Domain, error) {
	return d.conn.ListAllDomains()
}
