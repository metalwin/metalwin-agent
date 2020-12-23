package client

import (
	"strings"

	libvirt "github.com/metalwin/metalwin-agent/internal"
	libvirtxml "libvirt.org/libvirt-go-xml"
)

type domainService struct {
	conn libvirt.LibvirtClient
}

func (d *domainService) List() ([]libvirt.Domain, error) {
	return d.conn.ListAllDomains()
}

func (d *domainService) ListRaw() ([]libvirtxml.Domain, error) {
	domains, err := d.conn.ListAllDomains()
	if err != nil {
		return nil, err
	}
	if len(domains) == 0 {
		return make([]libvirtxml.Domain, 0), nil
	}
	var result []libvirtxml.Domain
	for _, domain := range domains {
		xml, err := domain.GetXMLDesc(0)
		if err != nil {
			continue
		}
		if xml == "" {
			continue
		}
		rawDomain := &libvirtxml.Domain{}
		err = rawDomain.Unmarshal(strings.TrimSpace(xml))
		if err != nil {
			continue
		}
		result = append(result, *rawDomain)
	}
	return result, nil
}
