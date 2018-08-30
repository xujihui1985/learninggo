package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c := make(chan interface{}, 0)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Printf("send %d to chan \n", i)
		}
		close(c)
	}()

	go func() {
		for {
			select {
			case j, ok := <-c:
				if !ok {
					cancel()
				}
				time.Sleep(1 * time.Second)
				fmt.Printf("process item from chan %d \n", j)
			}
		}
	}()
	<-ctx.Done()
}
