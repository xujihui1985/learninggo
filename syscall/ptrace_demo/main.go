package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	var regs syscall.PtraceRegs
	var ss syscallCounter

	ss = ss.init()

	fmt.Printf("Run %v\n", os.Args[1:])

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	cmd.Start()
	err := cmd.Wait()

	if err != nil {
		fmt.Printf("Wait returned: %v\n", err)
	}

	pid := cmd.Process.Pid
	err = syscall.PtraceGetRegs(pid, &regs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", regs)

	name := ss.getName(reg.Orig_rax)
	fmt.Printf("%s\n", name)
}
