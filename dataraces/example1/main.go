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
var counter int

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
		value := counter

		// yield the thread and be placed back in queue
		runtime.Gosched()

		// inc local value by 1
		value++

		// write back to global counter
		counter = value
	}
}

/*
==================
WARNING: DATA RACE
Read at 0x0000011d4318 by goroutine 7:
  main.incCounter()
      /Users/sean/code/go/src/github.com/xujihui1985/learninggo/dataraces/example1/main.go:34 +0x47
  main.main.func1()
      /Users/sean/code/go/src/github.com/xujihui1985/learninggo/dataraces/example1/main.go:21 +0x2f

Previous write at 0x0000011d4318 by goroutine 6:
  main.incCounter()
      /Users/sean/code/go/src/github.com/xujihui1985/learninggo/dataraces/example1/main.go:43 +0x68
  main.main.func1()
      /Users/sean/code/go/src/github.com/xujihui1985/learninggo/dataraces/example1/main.go:21 +0x2f

Goroutine 7 (running) created at:
  main.main()
      /Users/sean/code/go/src/github.com/xujihui1985/learninggo/dataraces/example1/main.go:20 +0xa2

Goroutine 6 (finished) created at:
  main.main()
      /Users/sean/code/go/src/github.com/xujihui1985/learninggo/dataraces/example1/main.go:20 +0xa2
==================
final counter: 4
Found 1 data race(s)
exit status 66
*/
