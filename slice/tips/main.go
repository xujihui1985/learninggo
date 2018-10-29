package main

import (
	"fmt"
)

func main() {
	// cloneSlice()
	// deleteElement()
	// deleteElement2()
	fmt.Println("insertSliceIntoAnother")
	insertSliceIntoAnother()

	fmt.Println("unshift")
	unshift()

	fmt.Println("popfront")
	popfront()
	fmt.Println("popback")
	popBack()

	s := []string{"a", "d", "c", "d", "b"}

	res := DeleteElements(s, func(s string) bool {
		return s != "d"
	}, true)
	fmt.Println(res)
	fmt.Println(s)
}

type Person struct {
	name string
	age  int
}

func cloneSlice() {
	var s []Person = nil
	s2 := s[:]
	sClone := append(s[:0:0], s...)
	fmt.Println(sClone == nil)
	fmt.Println(s2)
}

func deleteElement() {
	s := []string{"a", "b", "c", "d"}
	s = append(s[:1], s[2:]...)
	fmt.Println(s)
}

func deleteElement2() {
	s := []string{"a", "b", "c", "d"}
	i := 1
	fmt.Println(s[i+1:])
	res := copy(s[i:], s[i+1:])
	fmt.Println(res)
	fmt.Println(s)
	fmt.Println(s[:3])
}

func DeleteElements(s []string, shouldKeep func(string) bool, clear bool) []string {
	result := make([]string, 0, len(s))
	for _, v := range s {
		if shouldKeep(v) {
			result = append(result, v)
		}
	}
	if clear { // avoid memory leaking
		temp := s[len(result):]
		for i := range temp {
			temp[i] = "" // t0 is a zero value literal of T.
		}
	}
	return result
}

func insertSliceIntoAnother() {
	s := []string{"a", "b", "c", "d"}

	i := 2 // insertion position is 2

	e := []string{"e", "f", "g"}

	s = append(s[:i], append(e, s[i:]...)...)

	fmt.Println(s)
}

func push() {
	s := []string{"a", "b", "c", "d"}
	s = append(s, []string{"e", "f"}...)
}

func pushfront() {
	s := []string{"a", "b", "c", "d"}
	f := "e"
	s = append([]string{f}, s...)
	fmt.Println(s)
}

func unshift() {
	s := []string{"a", "b", "c", "d"}
	elements := []string{"e", "f"}
	s = append(elements, s...)
	fmt.Println(s)
}

func popfront() {
	s := []string{"a", "b", "c", "d"}
	s, e := s[1:], s[0]
	fmt.Println(e)
	fmt.Println(s)
}

func popBack() {
	s := []string{"a", "b", "c", "d"}
	s, e := s[:len(s)-1], s[len(s)-1]
	fmt.Println(e)
	fmt.Println(s)
}

func modifyArray() {
	[]int{1, 2, 3}[1] = 9
}
