package main

import "fmt"

type CustomSet map[string]string

func main() {

	s := CustomSet{
		"hello": "world",
	}

	for k, v := range s {
		fmt.Println(k, v)
	}

}
