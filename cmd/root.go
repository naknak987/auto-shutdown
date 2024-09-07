package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "auto_shutdown_server",
	Short: "Automatically shutdown your server",
	Long: `Automatically shutdown your server using ping
to detect a power failure.

Supports Proxmox and bare metal linux servers.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}