package debug

import (
	"log"
	"net/http"

	// debug pprof
	_ "expvar"
	_ "net/http/pprof"
	"time"
)

func Run() {
	debugHost := "localhost:8000"
	debug := http.Server{
		Addr:         debugHost,
		Handler:      http.DefaultServeMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Printf("Debug listening %s", debugHost)
		err := debug.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
