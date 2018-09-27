package main

import (
	"context"
	"fmt"
)

type user struct {
	name string
}

type userKey int

func main() {
	u := user{
		name: "sean",
	}

	const uk userKey = 0

	// Store the pointer to the user value inside the context
	// with a value of zero of type userkey
	ctx := context.WithValue(context.Background(), uk, &u)

	// retrieve that user pointer back by user the same key
	// type value
	if u, ok := ctx.Value(uk).(*user); ok {
		fmt.Println("User name", u.name)
	}

	// when retrieve the value using the same value but a different type
	if _, ok := ctx.Value(0).(*user); !ok {
		fmt.Println("user not found")
	}

}

// rule number 1
// ask again, do you really want to use context to store data
// rule number 2
// do not use primary type as key, define your own type, and use your own type as key to retrive value from context
