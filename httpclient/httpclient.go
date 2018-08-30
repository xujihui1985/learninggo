package main

import (
	"net/http"
	"time"
)

func main() {
	_ = &http.Client{
		Timeout:   time.Duration(3) * time.Second,
		Transport: http.DefaultTransport,
	}

}
