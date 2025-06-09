package secureexec

import (
	"context"
	"os/exec"
	"time"
)

func Command(name string, args ...string) *exec.Cmd {
	return exec.Command(name, args...)
}

func CommandWithTimeout(timeout time.Duration, name string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, name, args...)
	return cmd.CombinedOutput()
}
