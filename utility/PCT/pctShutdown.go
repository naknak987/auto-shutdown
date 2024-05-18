package PCT

// Shutdown methods for PCT machines (Proxmox CT containers)

import (
	"bytes"
	"os/exec"
	"errors"
)

type ctStat struct {
	vmID string
	status string
	lock string
	name string
}

func ShutdownAll() ([]string, error) {
	var ret []string

	cmd := exec.Command("pct", "list")
	stdOut, err := cmd.Output()
	if err != nil {
		return ret, err
	}

	lines := bytes.Split(stdOut, []byte("\n"))
	if len(lines) == 0 {
		return ret, errors.New("no ouput from command `pct list`")
	}

	for i, ln := range lines {
		// Skip the first line, header.
		if i == 0 {
			continue
		}

		// Break the line into individual fields.
		fields := bytes.Fields(ln)

		var pctResult ctStat
		fieldCount := len(fields)
		if fieldCount != 3 && fieldCount != 4 {
			return ret, errors.New("unexpected number of fields in output of `pct list`")
		} else if fieldCount == 3 {
			pctResult = ctStat{
				vmID: string(fields[0]),
				status: string(fields[1]),
				name: string(fields[2]),
			}
		} else {
			pctResult = ctStat{
				vmID: string(fields[0]),
				status: string(fields[1]),
				lock: string(fields[2]),
				name: string(fields[3]),
			}
		}

		// Issue shutdown commands.
		shutdownCmd := exec.Command("pct", "shutdown", pctResult.vmID)
		stdOut, err := shutdownCmd.Output()
		if err != nil {
			return ret, err
		}

		ret = append(ret, string(stdOut))
	}

	return ret, nil
}