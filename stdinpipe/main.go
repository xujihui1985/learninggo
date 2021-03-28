package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"syscall"
)

func main() {
	sqlBin, err := exec.LookPath("sqlite3")
	if err != nil {
		log.Fatalf("sqlite3 not found")
	}
	td, err := ioutil.TempDir("", "tmpsqlit3")
	if err != nil {
		log.Fatalf("error tmpdir")
	}
	dbpath := filepath.Join(td, "test.db")
	log.Printf("dbpath %s \n", dbpath)
	cmd := exec.Command(sqlBin, dbpath)
	w, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, "DROP TABLE IF EXISTS Bin;")
	fmt.Fprintln(w, "CREATE TABLE Bin(Func varchar, Pkg varchar);")
	fmt.Fprintln(w, "BEGIN TRANSACTION;")
	fmt.Fprintf(w, "INSERT INTO Bin VALUES(%q, %q);\n", "testfn", "aaaaa")
	fmt.Fprintln(w, "END TRANSACTION;")

	w.Close()

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	if err := syscall.Exec(cmd.Path, cmd.Args, cmd.Env); err != nil {
		log.Fatal(err)
	}
}
