package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func ExecuteCmd(command string, args ...string) (string, string, error) {
	cmd := exec.Command(command, args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	cmd.Start()
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case err := <-done:
		if err != nil {
			return out.String(), stderr.String(), fmt.Errorf("execution failed: %w", err)
		}
	case <-time.After(3 * time.Second): // Timeout after 3 seconds
		cmd.Process.Kill()
		return "", "Execution timed out", fmt.Errorf("execution timed out")
	}

	return out.String(), stderr.String(), nil
}
