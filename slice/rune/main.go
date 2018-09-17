package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "中文, hello world"

	for i, r := range s {
		rl := utf8.RuneLen(r)
		offset := i + rl
		fmt.Printf("codepoint: %#6x; offset: %d, char %s\n", r, offset, s[i:i+rl])
	}

}
