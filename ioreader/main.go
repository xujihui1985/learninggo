package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

func main() {
	//	runtime.GOMAXPROCS(2)
	//	bytes_reader()
	// tee_reader()
	// multi_reader()
	multi_writer()
}

func bytes_reader() {
	f, err := os.Open("/Users/sean/Downloads/buildingapplicationonmesos.pdf")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(b)

	data := make([]byte, 10)
	_, err = r.Read(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("data %+v\n", data)

	// rewind to start
	_, err = r.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}

	target, err := os.Create("./save")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(target, r)
	if err != nil {
		panic(err)
	}
}

func tmp_file() {

	f, err := ioutil.TempFile("", "upload")
	if err != nil {
		panic(err)
	}
	defer func() {
		n := f.Name()
		f.Close()
		os.Remove(n)
	}()
	s, err := os.Open("/Users/sean/Downloads/buildingapplicationonmesos.pdf")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(f, s)
	if err != nil {
		panic(err)
	}
	f.Seek(0, io.SeekStart)
	processFileMeta(f)
	// rewind to head
	f.Seek(0, io.SeekStart)
	handleFile(f)
}

func multi_reader() {
	f, err := os.Open("/Users/sean/Downloads/buildingapplicationonmesos.pdf")
	if err != nil {
		panic(err)
	}
	b := make([]byte, 2)
	_, err = f.Read(b)
	if err != nil {
		panic(err)
	}
	r := io.MultiReader(bytes.NewReader(b), f)
	handleFile(r)
}

func tee_reader() {
	f, err := os.Open("/Users/sean/Downloads/buildingapplicationonmesos.pdf")
	if err != nil {
		panic(err)
	}

	// create a pipe and tee reader
	pr, pw := io.Pipe()
	tr := io.TeeReader(f, pw)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer pw.Close()
		defer wg.Done()
		handleFile(tr)
	}()

	go func() {
		defer wg.Done()
		processFileMeta(pr)
	}()

	wg.Wait()
}

func multi_writer() {
	f, err := os.Open("/Users/sean/Downloads/buildingapplicationonmesos.pdf")
	if err != nil {
		panic(err)
	}
	ar, aw := io.Pipe()
	br, bw := io.Pipe()
	cr, cw := io.Pipe()
	dr, dw := io.Pipe()

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		handleFile(ar)
	}()

	go func() {
		defer wg.Done()
		handleFile1(br)
	}()
	go func() {
		defer wg.Done()
		handleFile2(cr)
	}()
	go func() {
		defer wg.Done()
		handleFile3(dr)
	}()

	go func() {
		defer func() {
			aw.Close()
			bw.Close()
			cw.Close()
			dw.Close()
		}()

		mw := io.MultiWriter(aw, bw, cw, dw)
		io.Copy(mw, f)
	}()

	wg.Wait()

}

func handleFile(r io.Reader) {
	target, err := os.Create("./save")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(target, r)
	if err != nil {
		panic(err)
	}
}

func handleFile1(r io.Reader) {
	target, err := os.Create("./save1")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(target, r)
	if err != nil {
		panic(err)
	}
}
func handleFile2(r io.Reader) {
	target, err := os.Create("./save2")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(target, r)
	if err != nil {
		panic(err)
	}
}
func handleFile3(r io.Reader) {
	target, err := os.Create("./save3")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(target, r)
	if err != nil {
		panic(err)
	}
}

func processFileMeta(r io.Reader) {
	target, err := os.Create("./save_meta")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(target, r)
	if err != nil {
		panic(err)
	}
}
