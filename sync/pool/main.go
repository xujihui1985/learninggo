package main

import (
	"fmt"
	"sync"
)

// const MaxPacketSize = 4096

var bufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 2)
	},
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	b := bufPool.Get().([]byte)
	fmt.Println(len(b), cap(b))
	b[0] = 1
	b[1] = 2
}
