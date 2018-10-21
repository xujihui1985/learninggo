package subbench

import (
	"fmt"
	"testing"
)

func BenchmarkSprint(b *testing.B) {
	b.Run("sprint", benchSprint)
	b.Run("sprintf", benchSprintf)
}

var gs string

func benchSprint(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello")
	}
	gs = s
}

func benchSprintf(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello")
	}
	gs = s
}
