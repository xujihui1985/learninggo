package main

import "fmt"

type u struct {
	name string
}

func main() {
	s := []u{
		u{
			name: "sean",
		},
	}
	for _, v := range s {
		(&v).name = "aaa"
		fmt.Println(v)
	}
	fmt.Println(s)

}
