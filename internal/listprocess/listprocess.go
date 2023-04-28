package listprocess

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/process"
)

func HasProcessRunning() {
	processes, _ := process.Processes()
	for _, p := range processes {
		fmt.Println(p.Name())
	}
}
