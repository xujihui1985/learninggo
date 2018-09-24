The operating system's scheduler is considered a `preemptive` scheduler

every time go application start up, it will look to see how many cores are available to it
if there's one then what's gonna happen is Go is gonna create what we call a logical processor


### GRQ

Global Run Queue These would be goroutines that are in a runnable state, but haven't been assigned to some P yet 

### LRQ

Local Run Queue

in Go, any function or method can be created to be a goroutine

when we start our Go program, the first thing the run time's gonna do is create that main goroutine and it's gonna put that in some Local Run Queue for some P. P is logic process

there are `four major places` in your code where the scheduler has an opportunity to make a scheduling decision

- `go` keyword which is what we are gonna use to create goroutines
- system call
- channel operations
- gc


Image we have a multi-threaded piece of software, this program launch two threads, T1 and T2, there will be lots of context switch here

when we use goroutine, from the operating system perspective, this thread never get switch, and it can get it's full time slice all the time


#### start up and shut down cleanly

