package main

import (
	"fmt"
	"reflect"
)

func main() {
	one, two, three := 0.1, 0.2, 0.3
	fmt.Printf("%t\n", one)
	fmt.Println(reflect.TypeOf(one))
	fmt.Println(one+two > three)
}
