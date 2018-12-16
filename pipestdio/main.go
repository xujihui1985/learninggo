package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	r := bufio.NewReader(os.Stdin)
	var b [1024]byte
	for {
		n, err := r.Read(b[:])
		if err != nil {
			break
		}
		fmt.Print(string(b[0:n]))
	}

}
