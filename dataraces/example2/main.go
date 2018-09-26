package main

// we can use -race to detect race
// go build -race
// go run -race xxxx

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// counter is a shared variable
var counter int64

func main() {

	const g = 2

	var wg sync.WaitGroup
	wg.Add(g)

	for i := 0; i < g; i++ {
		go func() {
			incCounter()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("final counter:", counter)

}

func incCounter() {
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
