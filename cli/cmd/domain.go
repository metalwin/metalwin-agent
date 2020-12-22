package cmd

import (
	"fmt"

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
		listDomains(c)
	},
}

func init() {
	domainCmd.AddCommand(domainLsCmd)
}

func listDomains(c *client.Client) {
	domains, _ := c.Domain.ListAll()
	for _, domain := range domains {
		name, _ := domain.GetName()
		fmt.Println(name)
	}
}
