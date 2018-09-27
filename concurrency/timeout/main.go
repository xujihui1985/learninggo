package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout := 100 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	// if use unbuffer channel here, there will be a leak, guess why?
	ch := make(chan struct{}, 1)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- struct{}{}
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	select {
	case <-ch:
		fmt.Println("work done")
	case <-ctx.Done():
		err := ctx.Err()
		if err == context.DeadlineExceeded {
			fmt.Println("process deadline exceed")
		}
		if err == context.Canceled {
			fmt.Println("process has been canceled")
		}
	}
}
