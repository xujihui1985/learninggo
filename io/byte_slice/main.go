package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("missing file name")
		os.Exit(1)
	}

	fileName := os.Args[1]
	// writeFile(fileName)
	// readFull(fileName)
	// bufferreader(fileName)
	// readfileFromReader(fileName)
	bufferIO(fileName)

}

func writeFile(fileName string) {
	b := []byte("hello world")
	ioutil.WriteFile(fileName, b, 0644)
}

func readFull(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var b [1000]byte
	// if byte array is large than file, there will be ErrUnexpectedEOF
	if _, err := io.ReadFull(f, b[0:]); err != nil {
		if err == io.ErrUnexpectedEOF {
			fmt.Println(err)
		}
	} else {
		fmt.Println(string(b[:]))
	}
}

func bufferreader(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if err := scanner.Err(); err != nil {
			panic(err)
		}
		fmt.Println("#", line)
	}
}

func readfileFromReader(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	var buf [1024]byte
	total := 0
	for {
		n, err := f.Read(buf[:])
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		total = total + n
	}

	fmt.Println(total)
}

func bufferIO(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		}
		d := strings.Fields(line)
		fmt.Println(d)
	}
}
