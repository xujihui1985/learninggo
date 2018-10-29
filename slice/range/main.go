package main

import "fmt"

type u struct {
	name string
}

func main() {
	s := []u{
		{
			name: "sean",
		},
	}
	for _, v := range s {
		(&v).name = "aaa"
		fmt.Println(v)
	}
	fmt.Println(s)

	for range s {

	}

	// the element iteration variable is omitted
	for key := range s {
		fmt.Println(key)
	}

	for key, element := range s {
		fmt.Println(key, element)
	}

	type Person struct {
		name string
		age  int
	}

	persons := [...]Person{{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		fmt.Println("before", i, p)
		// This modification has no effects on the iteration,
		// for the ranged array is a copy of the persons array.
		persons[1].name = "Jack"
		fmt.Println("after", i, p)
		// This modification has not effects on the persons array,
		// for p is just a copy of a copy of one persons element.
		p.age = 31
	}
	fmt.Println("persons:", persons)

}
