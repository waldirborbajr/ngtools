package killprocess

import (
	"errors"
	"fmt"
	"localhost/ngtools/internal/listprocess"
	"localhost/ngtools/internal/secureexec"
)

func KillRunningProcess(processName string) error {
	procArgs := []string{"-c", "killall", processName}

	listprocess.HasProcessRunning()

	cmd := secureexec.Command("bash", procArgs...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return errors.New(fmt.Sprintf("%s - %s", err, string(output)))
	}

	return nil
}
