package main

import "fmt"

const (
	a1 = iota
	a2
	a3
)

const (
	b1 = iota << 1
	b2
	b3
)

func main() {
	fmt.Println(a1, a2, a3)
	fmt.Println(b1, b2, b3)
}
