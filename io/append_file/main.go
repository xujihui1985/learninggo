package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("missing file name")
		os.Exit(1)
	}

	// Base return the last part of filename
	//	fmt.Println(filepath.Base("/User/zzz/aaaaa.go"))
	fileName := os.Args[1]
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "%s\n", "hello")
}
