package main

// we can use -race to detect race
// go build -race
// go run -race xxxx

import (
	"fmt"
	"runtime"
	"sync"
)

// counter is a shared variable
var (
	counter int
	mu      sync.Mutex
)

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
		// read the global counter into local value
		mu.Lock()
		{
			value := counter

			// yield the thread and be placed back in queue
			runtime.Gosched()

			// inc local value by 1
			value++

			// write back to global counter
			counter = value
		}
		mu.Unlock()
	}
}
