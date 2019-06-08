package processes

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
)

func GetAllProcesses() {
	pids, _ := process.Pids()

	for pid := range pids {
		fmt.Println(pid)
	}
}
