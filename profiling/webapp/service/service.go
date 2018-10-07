package service

import (
	"context"
	"expvar"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type key int

const (
	requestIDKey key = 0
)

var req = expvar.NewInt("requests")

var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	req.Add(1)
	time.Sleep(100)
	fmt.Fprintf(w, "hello")
})

// Run run app server
func Run() {
	router := http.NewServeMux()
	router.Handle("/search", handler)

	host := "localhost:5000"
	readTimeout := 10 * time.Second
	writeTimeout := 31 * time.Second
	idleTimeout := 15 * time.Second
	server := &http.Server{
		Addr:         host,
		Handler:      tracing(nextRequestID)(router),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("could not gracefully shutdown server %v\n", err)
		}
	}()

	log.Println("listening on", host)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not listen on %s, %v", host, err)
	}
}

func nextRequestID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-ID", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

}
