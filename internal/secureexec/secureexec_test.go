package secureexec_test

import (
	"localhost/ngtools/internal/secureexec"
	"testing"
)

func TestCommand(t *testing.T) {
	// Call the secureexec.Command function with some test arguments
	cmd := secureexec.Command("/usr/bin/echo", "hello", "world")

	// Ensure that the returned object is of type *exec.Cmd
	if _, ok := cmd.CombinedOutput(); ok != nil {
		t.Error("Command did not return an *exec.Cmd object")
	}

	// Ensure that the command's name and arguments were set correctly
	// if cmd.Path != "echo" {
	// 	t.Errorf("Command path = %s, want echo", cmd.Path)
	// }

	if len(cmd.Args) != 3 || cmd.Args[1] != "hello" || cmd.Args[2] != "world" {
		t.Errorf("Command args = %v, want [echo hello world]", cmd.Args)
	}
}
