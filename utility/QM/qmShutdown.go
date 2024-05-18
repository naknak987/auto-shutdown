package QM
// Shutdown methods for QM machines (Proxmox VMs)

import (
	"bytes"
	"os/exec"
	"errors"
)

type vmStat struct {
	vmID string
	name string
	status string
	mem string
	bootDisk string
	PID string
}

func ShutdownAll() ([]string, error) {
	var ret []string

	cmd := exec.Command("qm", "list")
	stdOut, err := cmd.Output()
	if err != nil {
		return ret, err
	}

	lines := bytes.Split(stdOut, []byte("\n"))
	if len(lines) == 0 {
		return ret, errors.New("no output from command `qm list`")
	}

	for i, ln := range lines {
		// Skip the first line, header.
		if i == 0 {
			continue
		}

		// Break the lines into individual fields.
		fields := bytes.Fields(ln)

		var qmResult vmStat
		if len(fields) != 6 {
			return ret, errors.New("unexpected number of fields in output of `qm list`")
		}

		qmResult = vmStat{
			vmID: string(fields[0]),
			name: string(fields[1]),
			status: string(fields[2]),
			mem: string(fields[3]),
			bootDisk: string(fields[4]),
			PID: string(fields[5]),
		}

		// Issue shutdown commands.
		shutdownCmd := exec.Command("qm", "shutdown", qmResult.vmID)
		stdOut, err := shutdownCmd.Output()
		if err != nil {
			return ret, err
		}

		ret = append(ret, string(stdOut))
	}

	return ret, nil
}