package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	f, err := os.OpenFile("mmap.bin", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	n, err := f.WriteAt([]byte{byte(0)}, 1<<8)
	if err != nil {
		panic(err)
	}
	fmt.Println("n = ", n)
	data, err := syscall.Mmap(int(f.Fd()), 0, 1<<8, syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		panic(err)
	}

	f.Close()

	for i, v := range []byte("hello syscall") {
		data[i] = v
	}

	fmt.Println(data)

	syscall.Munmap(data)
}
