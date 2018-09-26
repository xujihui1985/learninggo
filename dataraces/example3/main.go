package main

// we can use -race to detect race
// go build -race
// go run -race xxxx

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var shutdown int64

func main() {

	const g = 2

	var wg sync.WaitGroup
	wg.Add(g)

	for i := 0; i < g; i++ {
		go func(i int) {
			doWork(i)
			wg.Done()
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("shutdown now")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()

}

func doWork(id int) {
	for {
		fmt.Println("do work", id)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Println("shut down worker", id)
			break
		}
	}
}
