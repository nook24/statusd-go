package main

import (
	"fmt"
	"github.com/nook24/statusd-go/processes"
	"github.com/nook24/statusd-go/service"
	"github.com/nook24/statusd-go/win32services"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {

	v, _ := mem.VirtualMemory()
	swap, _ := mem.SwapMemory();

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println("Memory:")
	fmt.Println(v)

	fmt.Println("SWAP:")
	fmt.Println(swap)

	fmt.Println("Host info:")
	h, _ := host.Info()
	fmt.Println(h);
	fmt.Println("CPU Info:")
	cpuInfo, _ := cpu.Info()
	fmt.Println(cpuInfo)

	fmt.Println("CPU time")
	cpuTimes, _ := cpu.Times(true)
	fmt.Println(cpuTimes)

	fmt.Println("CPU percentage")
	cpuPercent, _ := cpu.Percent(0, true)
	fmt.Println(cpuPercent)

	fmt.Println("CPU load")
	loadAvg, _ := load.Avg()
	fmt.Println(loadAvg)

	partitions, _ := disk.Partitions(true)

	fmt.Println("All disks:")
	fmt.Println(partitions)

	for _, partition := range partitions {
		fmt.Println("Disk Usage:")
		fmt.Println(disk.Usage(partition.Mountpoint))
		fmt.Println("IO Counters:")
		fmt.Println(disk.IOCounters())
	}

	fmt.Println("Windows Services:")
	win32services.GetWindowsServices()

	fmt.Println("Processes:")
	processes.GetAllProcesses();

	//	http.HandleFunc("/", sayHello)
	//	if err := http.ListenAndServe(":8080", nil); err != nil {
	//		panic(err)
	//  }

	var services [10]*service.Service
	var cancelTokens [10]chan bool

	wg := new(sync.WaitGroup)

	for i := 0; i < len(services); i++ {
		cancelTokens[i] = make(chan bool)
		services[i] = service.NewService(time.Duration(rand.Int63n(10)+1)*time.Second, fmt.Sprintf("run command %d", i))
	}

	for i, srv := range services {
		go srv.Enqueue(cancelTokens[i], wg)
	}

	dur, _ := time.ParseDuration("120s")
	time.Sleep(dur)

	for _, ct := range cancelTokens {
		ct <- true
	}
	wg.Wait()

	fmt.Println("Shutdown successful")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	w.Write([]byte(message))
}
