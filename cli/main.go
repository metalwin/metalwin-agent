package main

import (
	"fmt"

	"github.com/metalwin/metalwin-agent/pkg/client"
	"github.com/metalwin/metalwin-agent/pkg/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "metalwin-agent",
	Short: "MetalWin agent CLI",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of all components",
	Run: func(cmd *cobra.Command, args []string) {
		c := createClient()
		showVersions(c)
	},
}

var (
	driver string
	path   string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&driver, "driver", "d", "qemu", "Driver used to connect to the libvirt daemon")
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "system", "Path used when connect to the libvirt daemon")
	rootCmd.AddCommand(versionCmd)
}

func main() {
	rootCmd.Execute()
}

func createClient() *client.Client {
	config := config.ConnectionConfig{
		Driver: driver,
		Path:   path,
	}
	client, err := client.NewClient(config)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Unable to connect to libvirt daemon with config: %v\n", config)
	}
	return client
}

func showVersions(c *client.Client) {
	hypervisorVersion, err := c.HypervisorVersion()
	if err != nil {
		fmt.Printf("Unable to get hypervisor version: %v\n", err)
	}
	fmt.Printf("hypervisor v%d\n", hypervisorVersion)
}
