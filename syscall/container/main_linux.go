// +build linux

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("help")
	}

}

// if we need to change the hostname of new created namespace,
// we must fork twice
func run() {
	fmt.Printf("Running from main %v\n", os.Args[2:])
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		// use user ns enable you to do something with root privilege inside container
		// notice that at this moment you can not use cgroup along with NEWUSER flag
		// | syscall.CLONE_NEWUSER,
		// Credential: &syscall.Credential{Uid: 0, Gid 0},
		// UidMappings: []syscall.SysProcIDMap {
		// 	{ContainerID: 0, HostID: os.Getpid(), Size: 1}
		// },
		// GidMappings: []syscall.SysProcIDMap {
		// 	{ContainerID: 0, HostID: os.Getpid(), Size: 1}
		// },
	}

	cmd.Run()
}

func child() {
	fmt.Printf("Running inside container %v\n", os.Args[2:])

	cgroup()
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("container")))
	must(syscall.Chroot("/rootfs"))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	must(syscall.Mount("something", "mytemp", "tmpfs", 0, ""))
	must(cmd.Run())
	must(syscall.Unmount("proc", 0))
	must(syscall.Unmount("mytemp", 0))
}

func cgroup() {
	cgroups := "/sys/fs/cgroup"

	mem := filepath.Join(cgroups, "memory")
	myCgroup := filepath.Join(filepath.Join(mem, "sean"))
	os.Mkdir(myCgroup, 0755)

	must(ioutil.WriteFile(filepath.Join(myCgroup, "memory.limit_in_bytes"), []byte("999424"), 0700))
	must(ioutil.WriteFile(filepath.Join(myCgroup, "memory.memsw.limit_in_bytes"), []byte("999424"), 0700))
	must(ioutil.WriteFile(filepath.Join(myCgroup, "notify_on_release"), []byte("1"), 0700))

	pid := strconv.Itoa(os.Getpid())
	must(ioutil.WriteFile(filepath.Join(myCgroup, "cgroup.procs"), []byte(pid), 0700))

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
