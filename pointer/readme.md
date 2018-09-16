### stack

the stack memory in go routinue start out as `2k` compare with `2m` default stack size for a new thread on linux

> On Linux/x86-32, the default stack size for a new thread is 2 megabytes. Under the NPTL threading implementation

### stack frame

every function is given a stack frame to work with, and the size of stack frame is known at compile time

```
func main() {
  count := 10  // we allocate 4 bytes of memory on the stackframe of main go routinue
}
```

### stack grow

continuous stack

when stack is not enough, it will create a new stack, and copy all stack frame to the new stack, stack memory can not be shared between
go routinue