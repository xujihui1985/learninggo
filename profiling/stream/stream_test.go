package stream

import (
	"bytes"
	"testing"
)

func BenchmarkAlgOne(b *testing.B) {
	var output bytes.Buffer
	in := assembleInputStream()

	find := []byte("elvis")
	repl := []byte("Elvis")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		algOne(in, find, repl, &output)
	}
}

func BenchmarkAlgTwo(b *testing.B) {

	var output bytes.Buffer
	in := assembleInputStream()

	find := []byte("elvis")
	repl := []byte("Elvis")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		algTwo(in, find, repl, &output)
	}
}

func BenchmarkAlgThree(b *testing.B) {

	var output bytes.Buffer
	in := assembleInputStream()

	find := []byte("elvis")
	repl := []byte("Elvis")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		algThree(in, find, repl, &output)
	}
}
