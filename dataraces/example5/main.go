package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	data      []string
	rw        sync.RWMutex
	readCount int64
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

}

func writer() {
	for i := 1; i <= 10; i++ {
		rw.Lock()
		{
			rc := atomic.LoadInt64(&readCount)
			fmt.Printf("Performing write: RCount[%d]\n", rc)
			data = append(data, fmt.Sprintf("String:%d", i))
		}
		rw.Unlock()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func reader(id int) {
	for {
		rw.RLock()
		{
			rc := atomic.AddInt64(&readCount, 1)
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			fmt.Printf("%d: Performing read: length[%d] RCount[%d]\n", id, len(data), rc)
			atomic.AddInt64(&readCount, -1)
		}
		rw.RUnlock()
	}
}
