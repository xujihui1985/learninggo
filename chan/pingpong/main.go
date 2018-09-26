package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	table := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		player("sean", table)
		wg.Done()
	}()

	go func() {
		player("anna", table)
		wg.Done()
	}()

	table <- 1
	wg.Wait()
	fmt.Println("game over")
}

func player(name string, table chan int) {
	for {
		ball, ok := <-table
		if !ok {
			fmt.Printf("player %s won\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("player %s missed\n", name)
			close(table)
			return
		}

		fmt.Printf("player %s hit %d\n", name, ball)
		ball++
		table <- ball
	}

}
