package killprocess

import (
	"errors"
	"fmt"
	"localhost/ngtools/internal/listprocess"
	"localhost/ngtools/internal/secureexec"
	"os"
)

func KillRunningProcess(processName string) error {
	if len(processName) == 0 {
		return errors.New("ERROR: missing process name")
	}

	procArgs := "killall -v " + processName

	if listprocess.HasProcessRunning() {
		cmd := secureexec.Command("sh", "-c", procArgs)
		cmd.Env = os.Environ()
		output, err := cmd.CombinedOutput()
		if err != nil {
			// return errors.New(fmt.Sprintf("%s - %s", err, string(output)))
			return (fmt.Errorf(string(output), err))
		}
	}

	return nil
}
