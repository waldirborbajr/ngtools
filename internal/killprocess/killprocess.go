package killprocess

import (
	"fmt"
	"localhost/ngtools/internal/listprocess"
	"localhost/ngtools/internal/secureexec"
	"os"
	"strconv"
)

func KillRunningProcess(processName string) error {
	if len(processName) == 0 {
		return fmt.Errorf("ERROR: missing process name")
	}

	procArgs := "killall -v " + processName

	if listprocess.HasProcessRunning() {
		cmd := secureexec.Command("sh", "-c", procArgs)
		cmd.Env = os.Environ()
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf(string(output), err)
		}
	}

	return nil
}

func KillNgrokProcess() error {
	data, err := os.ReadFile("ngrok.pid")
	if err != nil {
		return err
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return proc.Kill()
}
