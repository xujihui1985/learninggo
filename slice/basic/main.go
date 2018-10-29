package main

import "fmt"

type language struct {
	name string
	year int
}

type LangCategory struct {
	dynamic bool
	strong  bool
}

func main() {

	// Nested Composite Literals Can Be Simplified
	l := [...]language{
		{"C", 1972},
		{"Python", 1991},
	}

	fmt.Println(l)

	m := map[LangCategory]map[string]int{
		{true, true}: {
			"python": 1991,
			"Erlang": 1986,
		},
	}

	fmt.Println(m)

	var s []string

	for i, n := range s {
		fmt.Println(i, n)

	}
	fmt.Println(s)
	fmt.Println(s == nil)

	var m2 map[int][]string

	fmt.Println(m2)
	fmt.Println(m2 == nil)
	for _, i := range m2 {
		fmt.Println(i)
	}

}
