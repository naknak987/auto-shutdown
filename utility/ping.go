package utility

import (
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

func SinglePing(ip string) (int, error) {
	pinger, err := probing.NewPinger(ip)
	if err != nil {
		return 0, err
	}
	// Allow privileged ping, see readme.
	pinger.SetPrivileged(true)
	pinger.Timeout, err = time.ParseDuration("5s")
	if err != nil {
		return 0, err
	}
	pinger.ResolveTimeout = pinger.Timeout
	pinger.Count = 1
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		return 0, err
	}
	stats := pinger.Statistics()
	return stats.PacketsRecv, nil
}