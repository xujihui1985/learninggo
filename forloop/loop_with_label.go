package forloop

import "fmt"

func withLabel() {
	items := []string{"a", "b", "c", "d"}
	values := map[string][]string{
		"a": []string{"a", "aa", "aaa", "aaaa"},
		"b": []string{"b", "bb", "bbb", "bbbb"},
	}
Outer:
	for _, i := range items {
		fmt.Println("outer", i)
		for _, n := range values[i] {
			fmt.Println("inner", n)
			if n == "aa" {
				continue Outer
			}
		}
	}
}

func withIndex() {
	items := []string{"a", "b", "c", "d"}
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i])
	}
}
