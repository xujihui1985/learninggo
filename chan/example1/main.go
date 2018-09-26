package main

import (
	"fmt"
)

func main() {
	basic()

}

func basic() {
	ch := make(chan string)
	go func() {
		ch <- "hello"
	}()

	fmt.Println(<-ch)
}
