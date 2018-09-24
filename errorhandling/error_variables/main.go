package main

import (
	"errors"
	"fmt"
)

var (
	// the error variable convension is Errxxxx
	ErrBadRequest       = errors.New("Bad requests")
	ErrMovedPermanently = errors.New("moved permanently")
)

func main() {
	if err := call(true); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("bad request")
			return
		case ErrMovedPermanently:
			fmt.Println("moved permanently")
			return
		default:
			fmt.Println("unknown")
		}
	}

}

func call(v bool) error {
	if v {
		return ErrBadRequest
	}
	return ErrMovedPermanently
}
