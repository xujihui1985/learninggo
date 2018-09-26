package main

import (
	"fmt"
	"time"
)

// signal without data
func main() {
	signalClose()
}

func signalClose() {
	ch := make(chan struct{})
	fmt.Println("create channel")
	go func() {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("signal event")
		close(ch)
	}()

	fmt.Println("wait to be signaled")
	<-ch
	fmt.Println("event received")
}
