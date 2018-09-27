package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Worker interface {
	Work()
}

// Task provides a pool of goroutine that can execute any worker
type Task struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxG int) *Task {
	t := Task{
		// Using an unbuffered channel because we want the
		// guarantee of knowing the work being submitted is
		// actually being worked on after the call to run returns
		work: make(chan Worker),
	}
	t.wg.Add(maxG)
	for i := 0; i < maxG; i++ {
		go func() {
			for w := range t.work {
				w.Work()
			}
			t.wg.Done()
		}()
	}
	return &t
}

func (t *Task) Do(w Worker) {
	// we can use ctx to perform timeout or cancel
	// we can also record the counter of waiting worker to do
	t.work <- w
}

func (t *Task) Shutdown() {
	close(t.work)
	t.wg.Wait()
}

type logWriter struct {
	count int
}

func (l *logWriter) Work() {
	log.Println("print count", l.count)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
}

func main() {

	t := New(10)
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		l := logWriter{
			count: i,
		}
		t.Do(&l)
		wg.Done()
	}
	wg.Wait()

	// shutdown the task pool
	t.Shutdown()

}
