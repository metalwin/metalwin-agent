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
		showVersions(c)
	},
}

func showVersions(c *client.Client) {
	hypervisorVersion, err := c.HypervisorVersion()
	if err != nil {
		fmt.Printf("Unable to get hypervisor version: %v\n", err)
	}
	fmt.Printf("hypervisor v%d\n", hypervisorVersion)
}
