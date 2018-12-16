package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	PS, err := exec.LookPath("ps")
	if err != nil {
		panic(err)
	}

	env := os.Environ()
	err = syscall.Exec(PS, []string{"ps", "-a"}, env)
	panic(err)
}
