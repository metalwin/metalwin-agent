package cmd

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

var (
	driver string
	path   string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&driver, "driver", "d", "qemu", "Driver used to connect to the libvirt daemon")
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "system", "Path used when connect to the libvirt daemon")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(domainCmd)
}

// Execute the CLI
func Execute() error {
	return rootCmd.Execute()
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
