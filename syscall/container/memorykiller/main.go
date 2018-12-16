package main

import "fmt"

const kb = 1 << 20

func main() {
	var b [][]byte

	for i := 0; i < 1000; i++ {
		var k [kb]byte
		b = append(b, k[:])
		fmt.Printf("allocate %d kb\n", i+1)
	}

}
