package cmd

import (
	"fmt"

	"github.com/metalwin/metalwin-agent/pkg/client"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of all components",
	Run: func(cmd *cobra.Command, args []string) {
		c := createClient()
		defer c.Close()
		showVersions(c)
	},
}

func showVersions(c *client.Client) {
	hypervisorVersion, _ := c.HypervisorVersion()
	fmt.Printf("hypervisor: %s\n", hypervisorVersion)

	libvirtdVersion, _ := c.LibvirtDaemonVersion()
	fmt.Printf("libvirtd  : %s\n", libvirtdVersion)
}
