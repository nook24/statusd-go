package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var services [10]*Service
	var cancelTokens [10]chan bool

	wg := new(sync.WaitGroup)

	for i := 0; i < len(services); i++ {
		cancelTokens[i] = make(chan bool)
		services[i] = NewService(time.Duration(rand.Int63n(10)+1)*time.Second, fmt.Sprintf("run command %d", i))
	}

	for i, srv := range services {
		go srv.Enqueue(cancelTokens[i], wg)
	}

	dur, _ := time.ParseDuration("120s")
	time.Sleep(dur)

	for _, ct := range cancelTokens {
		ct<-true
	}
	wg.Wait()

	fmt.Println("Shutdown successful")
}

