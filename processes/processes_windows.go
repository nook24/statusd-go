package processes

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
)

func GetAllProcesses() {
	pids, _ := process.Pids()

	for pid := range pids {
		win32Proc, _ := process.GetWin32Proc(int32(pid))
		fmt.Println(win32Proc)
	}
}
