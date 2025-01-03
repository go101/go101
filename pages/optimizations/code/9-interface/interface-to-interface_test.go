package interfaces

import (
	"fmt"
	"io"
	"testing"
)

var v = 9999999
var x, y interface{}

func Benchmark_BoxBox(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = v // need one allocation
		y = v // need one allocation
	}
}

func Benchmark_BoxAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = v // need one allocation
		y = x // no allocations
	}
}

var s = "hello"

func Benchmark_BoxN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprint(io.Discard, s, s, s)
	}
}

func Benchmark_AssignN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var i interface{} = s
		fmt.Fprint(io.Discard, i, i, i)
	}
}
