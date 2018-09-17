package main

import "fmt"

func main() {
	case1()
	case2()
}

// 当capacity
func case1() {
	s1 := make([]string, 1, 20)
	s1[0] = "hello"
	p1 := &s1[0]
	s1 = append(s1, "world")
	*p1 = "hello2"
	fmt.Printf("value of p1 is %s, value of s1[0] is %s \n", *p1, s1[0])
}

func case2() {
	s1 := make([]string, 1)
	s1[0] = "hello"
	p1 := &s1[0]
	s1 = append(s1, "world")
	*p1 = "hello2"
	fmt.Printf("value of p1 is %s, value of s1[0] is %s \n", *p1, s1[0])
}