package main

import (
	"expvar"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/xujihui1985/learninggo/profiling/webapp/debug"
	"github.com/xujihui1985/learninggo/profiling/webapp/service"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

// expvars is adding the goroutine counts to the variable set.
func expvars() {
	gr := expvar.NewInt("goroutines")
	go func() {
		for range time.Tick(time.Millisecond * 250) {
			gr.Set(int64(runtime.NumGoroutine()))
		}
	}()
}

func main() {
	expvars()
	debug.Run()
	service.Run()
}
