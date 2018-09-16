package main

import "fmt"

type User struct {
	name string
}

func main() {
	u := escape()
	fmt.Println(u.name)
}

// when return address of user from stackframe
// escape analisys determine that User should live on heap
// instead of stack
func escape() *User {
	u := User{
		name: "sean",
	}
	return &u
}

// because return value are copied to upper stackframe (main function)
// so this stack frame is safe to be removed
func liveOnStack() User {
	u := User{
		name: "sean",
	}
	return u
}
