package utility

import (
	"fmt"
	"testing"
)

// Assuming DetectPVE() function lives in a utility package
func TestDetectPVE(t *testing.T) {
	t.Run("test detect PVE", func(t *testing.T) {
		result, err := DetectPVE()
		if err != nil {
			t.Errorf("DetectPVE() error = %v, wantErr %v", err, true)
		}

		fmt.Printf("t: %v\n", result)
	})
}
