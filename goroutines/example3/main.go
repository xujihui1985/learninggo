package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(2)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Create G")

	go func() {
		printPrime("A")
		wg.Done()
	}()

	go func() {
		printPrime("B")
		wg.Done()
	}()

	wg.Wait()
}

func printPrime(prefix string) {

NEXT:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue NEXT
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
