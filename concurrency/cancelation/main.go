package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// cancel must be called, or there will be leak, cancel can be called multiple times
	defer cancel()

	go func() {
		<-time.After(2 * time.Second)
		fmt.Println("call cancel")
		cancel()
	}()

	fmt.Println("start to do work")
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("canceled")
	}
}
