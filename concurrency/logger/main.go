package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// logger provider support to throw log away if log wirtes start to timeout due to latency
type logger struct {
	write chan string //
	wg    sync.WaitGroup
}

func New(w io.Writer, size int) *logger {
	l := logger{
		write: make(chan string, size),
	}

	// wg to track the write goroutine, we can add more G here to pick logger from write channel
	l.wg.Add(1)

	// this is the write goroutine that performs the actual writes
	go func() {
		// range over channel, once the channel is close and flushed
		// the loop will terminated
		for d := range l.write {
			fmt.Fprintln(w, d)
		}

		// when write channel closed, mark the logger shutdown
		l.wg.Done()
	}()

	return &l
}

func (l *logger) Shutdown() {

	// close the write channel
	close(l.write)

	// wait for the write goroutine to terminate and all log was flushed
	l.wg.Wait()
}

func (l *logger) Write(data string) {

	select {
	case l.write <- data:
		// log has been send to write channel
	default:
		// log was dropped, or we can stash the log
		fmt.Println("drop log")
	}
}

type device struct {
	off bool
}

func (d *device) Write(p []byte) (n int, err error) {
	if d.off {
		// simulate network error
		time.Sleep(2 * time.Second)
	}
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	fmt.Printf("current pid is %d\n", os.Getpid())
	const grs = 10
	var d device

	l := New(&d, grs)

	for i := 0; i < grs; i++ {
		go func(id int) {
			for {
				l.Write(fmt.Sprintf("%d: log data, pid: %d", id, os.Getpid()))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	// we want to control the simulated disk blocking.
	// capture USR1 signals to toggle device issues

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGUSR1)

	for {
		<-sigChan
		d.off = !d.off
	}
}
