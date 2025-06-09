package killprocess_test

import (
	"errors"
	"localhost/ngtools/internal/killprocess"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"
)

func TestKillRunningProcess(t *testing.T) {
	// Start a dummy process (sleep for 60 seconds)
	cmd := exec.Command("sleep", "60")
	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start dummy process: %v", err)
	}
	defer func() {
		// Kill the dummy process when the test is done
		_ = cmd.Process.Kill()
		_ = os.Remove("ngrok.pid")
	}()

	// Write the PID to ngrok.pid
	pid := cmd.Process.Pid
	if err := os.WriteFile("ngrok.pid", []byte(strconv.Itoa(pid)), 0600); err != nil {
		t.Fatalf("Failed to write ngrok.pid: %v", err)
	}

	// Give the process a moment to start
	time.Sleep(100 * time.Millisecond)

	// Attempt to kill the process using KillRunningProcess
	err := killprocess.KillRunningProcess()
	if err != nil {
		t.Errorf("KillRunningProcess returned error: %v", err)
	}

	// Check if the process is still running
	if err := cmd.Process.Signal(os.Signal(0)); err == nil {
		t.Errorf("Process should have been killed, but is still running")
	}

	// Test with missing PID file
	_ = os.Remove("ngrok.pid")
	err = killprocess.KillRunningProcess()
	if err == nil {
		t.Errorf("Expected error when PID file is missing, got nil")
	}

	// Test with invalid PID in file
	if err := os.WriteFile("ngrok.pid", []byte("invalid"), 0600); err != nil {
		t.Fatalf("Failed to write invalid ngrok.pid: %v", err)
	}
	err = killprocess.KillRunningProcess()
	if err == nil {
		t.Errorf("Expected error with invalid PID, got nil")
	}

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
