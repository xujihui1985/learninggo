// always return error interface
package main

import "fmt"

type CustomError struct {
}

func (*CustomError) Error() string {
	return ""
}

func dosth() error {
	return nil
}

func dosth2() error {
	var e *CustomError
	return e
}

func main() {
	var err error = nil
	var err2 *CustomError = nil
	var err3 *CustomError = err2
	err4 := dosth()
	err5 := err2
	var err6 error = err2
	err7 := dosth2()

	fmt.Println("err == nil", err == nil)
	fmt.Println("err2 == nil", err2 == nil)
	fmt.Println("err3 == nil", err3 == nil)
	fmt.Println("err4 == nil", err4 == nil)
	fmt.Println("err5 == nil", err5 == nil)
	fmt.Println("err6 == nil", err6 == nil)
	fmt.Println("err7 == nil", err7 == nil)
}
