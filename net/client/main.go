package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {

	c, err := net.Dial("unix", "/tmp/test.sock")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	reader(c)
	for {
		r := bufio.NewReader(os.Stdin)
		fmt.Println("text to send")
		text, _ := r.ReadString('\n')
		fmt.Println(text)
		c.Write([]byte(text))

		bufio.NewReader(c)
		_, err = c.Write([]byte("hello"))
		if err != nil {
			panic(err)
		}
		time.Sleep(2 * time.Second)
	}
}

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		fmt.Println("read from conn")
		n, err := r.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("got data", string(buf[0:n]))
	}
}
