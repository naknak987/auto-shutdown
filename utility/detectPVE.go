package utility

import (
	"os/exec"
)

func DetectPVE() (bool, error) {
	// If the command does not produce an error, we are on ProxMox
	cmd := exec.Command("pveversion")
	_, err := cmd.Output()
	if err != nil {
		return false, err
	}

	return true, nil
}