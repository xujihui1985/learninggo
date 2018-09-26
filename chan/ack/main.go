package main

import (
	"fmt"
	"time"
)

func main() {
	// signalAck()
	// selectSend()
	// selectRecv()
	selectDrop()
}

func signalAck() {
	ch := make(chan string)
	go func() {
		fmt.Println(<-ch)
		time.Sleep(1 * time.Second)
		ch <- "done"
	}()

	time.Sleep(1 * time.Second)
	ch <- "start working"
	fmt.Println(<-ch)
}

func selectSend() {
	ch := make(chan string)
	go func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("wait for work")
		fmt.Println(<-ch)
	}()

	select {
	case ch <- "work":
		fmt.Println("send work")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out")
	}
}

func selectRecv() {
	ch := make(chan string)
	go func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("send to ch")
		// this send can not finish unless there is a corresponding receive
		ch <- "work"
		fmt.Println("after send to ch")
	}()

	select {
	case v := <-ch:
		fmt.Println("receive work", v)
		// if timeout happened before we receive data from ch, there is a leak
		// to go routine above can never terminate
	case <-time.After(3000 * time.Millisecond):
		fmt.Println("timed out")
	}

}

func selectDrop() {
	ch := make(chan int, 5)
	go func() {
		for v := range ch {
			fmt.Println("recv", v)
		}
	}()

	for i := 0; i < 100; i++ {
		select {
		case ch <- i:
		default:
			// when ch is block, it will drop the request,
			// image i is request coming from outside
			fmt.Println("Drop", i)
		}
	}
	close(ch)
}
