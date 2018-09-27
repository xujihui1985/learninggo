package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	deadline := time.Now().Add(150 * time.Millisecond)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	defer cancel()

	// if use unbuffer channel here, there will be a leak, guess why?
	ch := make(chan struct{}, 1)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		fmt.Println("work done")
	case <-ctx.Done():
		fmt.Printf("%v\n", ctx.Err())
	}
}
