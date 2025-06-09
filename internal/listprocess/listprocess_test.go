package listprocess

import (
    "os"
    "os/exec"
    "strconv"
    "testing"
    "time"
)

func TestHasProcessRunning(t *testing.T) {
    // Start a dummy process (sleep for 30 seconds)
    cmd := exec.Command("sleep", "30")
    if err := cmd.Start(); err != nil {
        t.Fatalf("Failed to start dummy process: %v", err)
    }
    defer func() {
        _ = cmd.Process.Kill()
    }()

    // Give the process a moment to start
    time.Sleep(100 * time.Millisecond)

    // Check if the process is running
    running, err := HasProcessRunning("sleep")
    if err != nil {
        t.Errorf("HasProcessRunning returned error: %v", err)
    }
    if !running {
        t.Errorf("Expected sleep process to be running, but it was not detected")
    }

    // Check for a process that does not exist
    running, err = HasProcessRunning("definitelynotarunningprocess")
    if err != nil {
        t.Errorf("HasProcessRunning returned error: %v", err)
    }
    if running {
        t.Errorf("Expected process to not be running, but it was detected")
    }
}

// Optional: Test HasProcessRunning with empty process name
func TestHasProcessRunningEmptyName(t *testing.T) {
    running, err := HasProcessRunning("")
    if err == nil {
        t.Errorf("Expected error for empty process name, got nil")
    }
    if running {
        t.Errorf("Expected running to be false for empty process name")
    }