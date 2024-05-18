package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var daemonTestCmd = &cobra.Command{
	Use:     "daemon_test",
	Aliases: []string{"dtest"},
	Short:   "Mimicks output from real Proxmox ",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "pct"{
			switch args[1] {
			case "list":
				pmtList()
			default:
				fmt.Println("No arguments for specified command")
			}
		} else {
			fmt.Println("No command specified")
		}
	},
}

func init() {
	rootCmd.AddCommand(daemonTestCmd)
}

func pmtList() {
	fmt.Printf(`VMID       Status     Lock         Name
101        stopped                 ai
103        stopped                 runner
104        running                 go-tester`)
}