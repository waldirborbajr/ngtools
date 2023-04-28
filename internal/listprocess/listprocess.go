package listprocess

import (
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

func HasProcessRunning() bool {
	processes, _ := process.Processes()

	for _, p := range processes {
		procName, _ := p.Name()
		if strings.Contains(procName, "ngrok") {
			return true
		}
	}
	return false
}
