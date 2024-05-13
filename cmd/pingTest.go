package cmd

import (
	"fmt"

	"gitea.naknak987.com/auto-shutdown-server/v2/utility"
	"github.com/spf13/cobra"
)

var pingTestCmd = &cobra.Command{
	Use:     "ping_test",
	Aliases: []string{"ping"},
	Short:   "Starts ping testing minutely against the IP address passed.",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result, _ := utility.SinglePing(args[0])
		fmt.Printf("Successful pings: %d", result)
	},
}

func init() {
	rootCmd.AddCommand(pingTestCmd)
}

