package cmd

import (
	"fmt"

	"gitea.naknak987.com/auto-shutdown-server/v2/utility"
	"gitea.naknak987.com/auto-shutdown-server/v2/utility/PCT"
	"github.com/spf13/cobra"
)

var randTestCmd = &cobra.Command{
	Use:     "rand_test",
	Aliases: []string{"test"},
	Short:   "Starts ping testing minutely against the IP address passed.",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ret, err := PCT.ShutdownAll()
		if err != nil {
			fmt.Println(err)
		}

		for i, v := range ret {
			fmt.Printf("%d %s\n", i, v)
		}
	},
}

func init() {
	rootCmd.AddCommand(randTestCmd)
}

func pingTest(ip string) {
	result, err := utility.SinglePing(ip)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successful pings: %d\n", result)
}
