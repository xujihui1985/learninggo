package main

const (
	//	size = 10

	size = 1024
)

func main() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

func stackCopy(s *string, c int, a [size]int) {
	println(c, s, *s)
	c++
	if c == 10 {
		return
	}
	stackCopy(s, c, a)
}

/**
### stack grow

continuous stack

when stack is not enough, it will create a new stack, and copy all stack frame to the new stack, stack memory can not be shared between
go routinue
**/

/**
when size is 10, the address of string is
0 0xc42005df68 HELLO
1 0xc42005df68 HELLO
2 0xc42005df68 HELLO
3 0xc42005df68 HELLO
4 0xc42005df68 HELLO
5 0xc42005df68 HELLO
6 0xc42005df68 HELLO
7 0xc42005df68 HELLO
8 0xc42005df68 HELLO
9 0xc42005df68 HELLO
**/

/**
when size is 1024, the address of string is
0 0xc42009df68 HELLO
1 0xc42009df68 HELLO
2 0xc4200adf68 HELLO
3 0xc4200adf68 HELLO
4 0xc4200adf68 HELLO
5 0xc4200adf68 HELLO
6 0xc4200cdf68 HELLO
7 0xc4200cdf68 HELLO
8 0xc4200cdf68 HELLO
9 0xc4200cdf68 HELLO
**/
