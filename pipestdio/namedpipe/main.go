package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	// createLogEntry()
	// appendLog()
	namedPipe()
}

func namedPipe() {
	namedPipe := filepath.Join("/var/applogs", "stdout")
	fmt.Println(namedPipe)
	syscall.Mkfifo(namedPipe, 0600)

	stdout, _ := os.OpenFile(namedPipe, os.O_RDONLY, 0600)
	defer stdout.Close()
	fmt.Println("reading")

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM, syscall.SIGPIPE)
	go func(c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		os.Exit(0)
	}(sigc)

	var b [512]byte
	httpClient := http.DefaultClient
	for {
		fmt.Println("reading from namedpipe")
		n, err := stdout.Read(b[:])
		if err != nil {
			fmt.Println(err)
			continue
		}
		req, _ := http.NewRequest("PUT", "https://tfapi.alipay.com/api/v1/logs/5be2f1890ab517886f7fff86", bytes.NewReader(b[0:n]))
		resp, err := httpClient.Do(req)
		if err != nil {
			panic(err)
		}
		resp.Body.Close()
	}
}

func createLogEntry() {
	httpClient := &http.Client{}
	// b, _ := json.Marshal(struct {
	// 	append bool
	// }{
	// 	append: true,
	// })

	req, _ := http.NewRequest("POST", "https://tfapi.alipay.com/api/v1/logs?append", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func appendLog() {
	httpClient := http.DefaultClient
	req, _ := http.NewRequest("PUT", "https://tfapi.alipay.com/api/v1/logs/5be2ddfff8b5f7fa3ff36831", bytes.NewReader([]byte("hello world")))
	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)

}
