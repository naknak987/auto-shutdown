package cmd

import (
	"time"

	"gitea.naknak987.com/auto-shutdown-server/v2/utility"
	"github.com/spf13/cobra"
	"github.com/coreos/go-systemd/v22/daemon"
	"github.com/coreos/go-systemd/v22/journal"
)

var startDaemonCmd = &cobra.Command{
	Use:     "start_daemon",
	Aliases: []string{"start"},
	Short:   "Starts a daemon that checks if this machine can reach another known machine on the network.",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		daemonStart(args[0])
	},
}

func init() {
	rootCmd.AddCommand(startDaemonCmd)
}

func runPing(ip string, failures *int) {
	res, err := utility.SinglePing(ip)
	if err != nil {
		// An error, log it to the systemd journal.
		journal.Send(err.Error(), journal.PriCrit, nil)
	}
	if res == 0 {
		// We had a failure, increment the failure counter.
		*failures += 1
	} else if *failures != 0 {
		// Reset failure counter to zero.
		*failures = 0
	}
}

func daemonStart(ip string) {
	// Setup
	failures := 0
	sleepDuration, err := time.ParseDuration("1m")
	if err != nil {
		// An error, log it to the systemd journal.
		journal.Send(err.Error(), journal.PriCrit, nil)
	}

	// Tell systemd we started successfully.
	daemon.SdNotify(false, daemon.SdNotifyReady)

	// Main loop
	for {
		// if we exceed 4 failures, shutdown the servers.
		if failures > 4 {
			journal.Send("5 failures, shutting down.", journal.PriAlert, nil)
			daemon.SdNotify(false, daemon.SdNotifyStopping)
			break
		}

		// Run ping test.
		go runPing(ip, &failures)

		// Wait for duration.
		time.Sleep(sleepDuration)
	}
}
