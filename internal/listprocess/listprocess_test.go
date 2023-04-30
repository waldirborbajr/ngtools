package listprocess_test

import (
	"fmt"
	"io/ioutil"
	"localhost/ngtools/internal/listprocess"
	"os"
	"os/exec"
	"syscall"
	"testing"
)

func TestHasProcessRunning(t *testing.T) {
	// Start a dummy process with "ngrok" in its name
	// to simulate a running ngrok process
	dummyProcess := startDummyProcess("ngrok process")

	// Check if the function correctly detects the running process
	if !listprocess.HasProcessRunning() {
		t.Errorf("HasProcessRunning() returned false, expected true")
	}

	// Kill the dummy process
	dummyProcess.Kill()

	// Check if the function correctly detects the stopped process
	if listprocess.HasProcessRunning() {
		t.Errorf("HasProcessRunning() returned true, expected false")
	}
}

// Helper function to start a dummy process for testing purposes
// func startDummyProcess(name string) *os.Process {
// 	cmd := exec.Command("sleep", "3600")
// 	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
// 	err := cmd.Start()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Rename the process to include the name passed as argument
// 	_, err = os.Stat(fmt.Sprintf("/proc/%d", cmd.Process.Pid))
// 	if err == nil {
// 		err = syscall.Prctl(syscall.PR_SET_NAME, []byte(name), 0, 0, 0)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	return cmd.Process
// }

// Helper function to start a dummy process for testing purposes
func startDummyProcess(name string) *os.Process {
	cmd := exec.Command("sleep", "3600")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// Modify the process name by writing to the /proc/self/comm file
	commPath := fmt.Sprintf("/proc/%d/comm", cmd.Process.Pid)
	err := ioutil.WriteFile(commPath, []byte(name), 0644)
	if err != nil {
		panic(err)
	}

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	return cmd.Process
}
