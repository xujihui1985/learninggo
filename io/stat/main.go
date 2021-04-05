package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fileName := os.Args[1]
	fi, err := os.Stat(fileName)
	if err != nil {
		panic(err)
	}

	if s, ok := fi.Sys().(*syscall.Stat_t); ok {
		fmt.Printf("%v\n", s)
	}

}
