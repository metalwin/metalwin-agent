package main

import (
	"log"

	"github.com/metalwin/metalwin-agent/pkg/client"
	"github.com/metalwin/metalwin-agent/pkg/config"
)

func main() {
	config := config.ConnectionConfig{
		Driver: "qemu",
		Path:   "system",
	}
	client, err := client.NewClient(config)
	if err != nil {
		log.Printf("Error creating libvir client: %s\n", err.Error())
		return
	}
	version, err := client.Version()
	if err != nil {
		log.Printf("Cannot retrieve hypervisor version: %s", err)
		return
	}
	log.Printf("Hypervisor version: %d", version)
}
