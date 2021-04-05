package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	var l net.Listener
	_, err := os.Stat("/tmp/test.sock")
	if os.IsNotExist(err) {
		l, err = net.Listen("unix", "/tmp/test.sock")
	}
	f, _ := l.(*net.UnixListener).File()
	l2, err := net.FileListener(f)
	if err != nil {
		panic(err)
	}
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM, syscall.SIGPIPE)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ln net.Listener, c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		ln.Close()
		os.Exit(0)
	}(l, sigc)
	go func() {
		for {
			fmt.Println("waiting for l2 to accept connection")
			conn, err := l2.Accept()
			if err != nil {
				fmt.Printf("failed to accept connection %+v\n", err)
				continue
			}
			fmt.Println("connection accepted from listener l2")
			go handleConn(conn)
		}
	}()
	go func() {
		n := 0
		for {
			if n = n + 1; n == 5 {
				break
			}
			fmt.Println("waiting for l1 to accept connection")
			conn, err := l.Accept()
			if err != nil {
				fmt.Printf("failed to accept connection %+v\n", err)
				continue
			}
			fmt.Println("connection accepted from listener l1")
			go handleConn(conn)
		}

	}()
	wg.Wait()
}

func handleConn(c net.Conn) {
	defer c.Close()
	net.Pipe()
	buf := make([]byte, 1024)
	for {
		n, err := c.Read(buf[0:])
		if err != nil {
			return
		}
		data := buf[0:n]
		fmt.Println(string(data))
		_, err = c.Write([]byte("write from server"))
		if err != nil {
			fmt.Println("write error: ", err)
		}
	}
}
