package secureexec

import (
	"os/exec"
)

// func Command(name string, args ...string) (*exec.Cmd, *strings.Builder) {
// 	return exec.Command(name, args...), new(strings.Builder)
// }

func Command(name string, args ...string) *exec.Cmd {
	return exec.Command(name, args...)
}
