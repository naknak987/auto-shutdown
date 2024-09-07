package cmd

import (
	"os/exec"
	"time"

	"github.com/naknak987/auto-shutdown/utility"
	"github.com/coreos/go-systemd/v22/daemon"
	"github.com/coreos/go-systemd/v22/journal"
	"github.com/spf13/cobra"
)

var MinutesWithoutPower int

var startDaemonCmd = &cobra.Command{
	Use:     "start_daemon [ip_address]",
	Aliases: []string{"start"},
	Short:   "Starts a daemon that checks if this machine can reach another known machine on the network.",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		daemonStart(args[0])
	},
}

func init() {
	startDaemonCmd.Flags().IntVarP(&MinutesWithoutPower, "minutes-without-power", "m", 5, "how many minutes without power before shutdown")
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
	journal.Send("Auto Shutdown Daemon started successfully", journal.PriInfo, nil)

	// Main loop
	for {
		// if we exceed (MinutesWithoutPower Int) failures, shutdown the servers.
		if failures > MinutesWithoutPower {
			journal.Send("5 failures, shutting down.", journal.PriNotice, nil)

			isPVE, err := utility.DetectPVE()
			if err != nil {
				journal.Send(err.Error(), journal.PriCrit, nil)
			} else if isPVE {
				shutdownPVE()
				time.Sleep(sleepDuration)
			}

			shutdownLinux();

			daemon.SdNotify(false, daemon.SdNotifyStopping)
			break;			
		}

		// Run ping test.
		go runPing(ip, &failures)

		// Wait for duration.
		time.Sleep(sleepDuration)
	}
}

func shutdownPVE() {
	// Shutdown containers
	result, err := utility.PCTShutdownAll()
	if err != nil {
		journal.Send(err.Error(), journal.PriCrit, nil)
	}
	for _, v := range result {
		journal.Send(v, journal.PriInfo, nil)
	}

	// Shutdown virtual machines
	result, err = utility.QMShutdownAll()
	if err != nil {
		journal.Send(err.Error(), journal.PriCrit, nil)
	}
	for _, v := range result {
		journal.Send(v, journal.PriInfo, nil)
	}
}

func shutdownLinux() {
	// Shutdown system
	journal.Send("Running system shutdown command", journal.PriInfo, nil)
	sdCmd := exec.Command("shutdown", "+1")
	sdRes, err := sdCmd.Output()
	if err != nil {
		journal.Send(err.Error(), journal.PriCrit, nil)
	}
	journal.Send(string(sdRes), journal.PriInfo, nil)
}