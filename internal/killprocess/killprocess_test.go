package killprocess_test

import (
	"errors"
	"localhost/ngtools/internal/killprocess"
	"os"
	"os/exec"
	"testing"
)

func TestKillRunningProcess(t *testing.T) {
	// Start a dummy process
	cmd := exec.Command("sleep", "3600")
	err := cmd.Start()
	if err != nil {
		t.Fatalf("Failed to start dummy process: %v", err)
	}
	defer func() {
		// Kill the dummy process when the test is done
		err := cmd.Process.Kill()
		if err != nil {
			t.Fatalf("Failed to kill dummy process: %v", err)
		}
	}()

	// Test case 1: Process name is empty
	err = killprocess.KillRunningProcess("")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	// Test case 2: Process name is valid, but no process is running
	err = killprocess.KillRunningProcess("notepad.exe")
	if err != nil {
		t.Errorf("Expected nil, but got %v", err)
	}

	// Test case 3: Process name is valid and process is running
	err = killprocess.KillRunningProcess("sleep")
	if err != nil {
		if !errors.Is(err, os.ErrPermission) {
			t.Errorf("Expected permission denied error, but got %v", err)
		}
	}
}
