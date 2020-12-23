package cmd

import (
	"fmt"
	"strings"

	"github.com/metalwin/metalwin-agent/pkg/client"
	"github.com/spf13/cobra"
)

var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Manage domains from the host",
}

var domainLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all domains",
	Run: func(cmd *cobra.Command, args []string) {
		c := createClient()
		defer c.Close()
		listDomains(c)
	},
}

func init() {
	domainCmd.AddCommand(domainLsCmd)
}

func listDomains(c *client.Client) {
	domains, _ := c.Domain.ListRaw()
	for _, domain := range domains {
		uuid := strings.TrimSpace(domain.UUID)
		fmt.Printf("%s\t%s\t%s\t%d %s\t%d vCpu\n",
			uuid, domain.Name, domain.OS.Type.Arch, domain.Memory.Value, domain.Memory.Unit, domain.VCPU.Value)
	}
}
