package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start go routines")

	// create a goroutine, this tells go process to not execute this function immeditately, load it in some local run queue for our P
	go func() {

		wg.Done()
	}()
	// now there are two G, one is main another is above

	// create another goroutine
	go func() {

		wg.Done()
	}()
	// now there are three G

	fmt.Println("wait to finish")
	wg.Wait()

	fmt.Println("terminate")
}
