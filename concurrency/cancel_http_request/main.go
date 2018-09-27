package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "https://yuque.antfin-inc.com/", nil)
	if err != nil {
		log.Println(err)
		return
	}

	var tr http.Transport
	client := http.Client{
		Transport: &tr,
	}

	ch := make(chan error, 1)

	go func() {
		log.Println("Starting request")
		resp, err := client.Do(req)
		if err != nil {
			ch <- err
			return
		}
		defer resp.Body.Close()
		io.Copy(os.Stdout, resp.Body)
		ch <- nil
	}()

Outer:
	for {
		select {
		case <-ctx.Done():
			tr.CancelRequest(req)
		case err := <-ch:
			if err != nil {
				log.Println(111, err)
			}
			break Outer
		}
	}

}
