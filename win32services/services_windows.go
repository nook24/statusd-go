package win32services

import (
	"fmt"
	"github.com/shirou/gopsutil/winservices"
)

func GetWindowsServices() {
	windowsServices, _ := winservices.ListServices()
	for _, winServices := range windowsServices {
		fmt.Println(winServices.Name)
		fmt.Println(winServices.Config)
		fmt.Println(winServices.Status)
	}
}
