package main

func main() {
	example(make([]string, 2, 2), "hello", 10)
}

//go:noinline
func example(slice []string, str string, i int) {
	panic("stack trace")
}

/*
panic: stack trace

goroutine 1 [running]:
main.example(0xc000034768, 0x2, 0x2, 0x106ab2e, 0x5, 0xa)
	/Users/sean/work/learninggo/profiling/stacktrace/example1/main.go:9 +0x39
main.main()
	/Users/sean/work/learninggo/profiling/stacktrace/example1/main.go:4 +0x68

the stack trace also shows the detail information of the parameter

main.example(0xc000034768, 0x2, 0x2, 0x106ab2e, 0x5, 0xa)

the first 3 params represent string slice, the first is the pointer to the back array
, the second hex 0x2 means length 2 and the third hex 0x2 means capacity 2

0xc000034768: the address of array the slice ref to
0x2: the length of slice
0x2: the capacity of slice
0x106ab2e: the address of byte array
0x5: the length of byte array
0xa: 10
*/
