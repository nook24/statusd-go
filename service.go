package main

import (
	"fmt"
	"sync"
	"time"
)

type Service struct {
	checkInterval time.Duration
	command string

	lastCheck time.Time
}

func NewService(checkInterval time.Duration, command string) *Service {
	s := new(Service)
	s.checkInterval = checkInterval
	s.command = command
	s.lastCheck = time.Time{}
	return s
}

func (s *Service) GetCheckInterval() time.Duration {
	return s.checkInterval
}

func (s *Service) GetCommand() string {
	return s.command
}

func (s *Service) GetLastCheck() time.Time {
	return s.lastCheck
}

func (s *Service) GetNextCheck() time.Time {
	return s.lastCheck.Add(s.checkInterval)
}

func (s *Service) Run() {
	fmt.Printf("Run some command %s\n", s.GetCommand())
	s.lastCheck = time.Now()
}

func (s *Service) Enqueue(cancel chan bool, wg *sync.WaitGroup) {
	for {
		t := time.NewTimer(s.GetNextCheck().Sub(time.Now()))

		select {
		case <-t.C:
			s.Run()
		case <-cancel:
			fmt.Println("Canceled")
			t.Stop()
			wg.Done()
			return
		}
	}
}
