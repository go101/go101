package interfaces

import "testing"

var r interface{}

var v0 struct{}

func Benchmark_BoxZeroSize1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = v0
	}
}

var a0 [0]int64

func Benchmark_BoxZeroSize2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = a0
	}
}

var b bool

func Benchmark_BoxBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = b
	}
}

var n int8 = -100

func Benchmark_BoxInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = n
	}
}
