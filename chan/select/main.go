package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

const timeout = 10 * time.Second

var (
	sigChan      = make(chan os.Signal, 1)
	timeoutChan  = time.After(timeout)
	completeChan = make(chan error)
	shutdownChan = make(chan struct{})
)

func main() {
	fmt.Println("starting process")
	signal.Notify(sigChan, os.Interrupt)
	go processor(completeChan)

ContinueLoop:
	for {
		select {
		case <-sigChan:
			fmt.Println("os interrupt")
			close(shutdownChan)
			// set the chan to nil so we no longer process these events
			sigChan = nil
		case <-timeoutChan:
			fmt.Println("timeout killing the program")
			os.Exit(1)
		case err := <-completeChan:
			fmt.Println("complete process", err)
			break ContinueLoop
		}
	}

}

// receive only chan
func processor(complete chan<- error) {
	var err error
	// defer function execute after processor function return
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("processor panic", r)
		}
		complete <- err
	}()
	err = doWork()
	if err != nil {
		return
	}
	fmt.Println("processor completed")
}

func doWork() error {
	fmt.Println("process task1")
	time.Sleep(2 * time.Second)
	if checkShutdown() {
		return errors.New("shutdown")
	}

	fmt.Println("process task2")
	time.Sleep(1 * time.Second)

	if checkShutdown() {
		return errors.New("shutdown")
	}

	fmt.Println("process task3")
	time.Sleep(1 * time.Second)
	return nil
}

func checkShutdown() bool {
	select {
	case <-shutdownChan:
		fmt.Println("process shutdown")
		return true
	default:
		return false
	}
}
