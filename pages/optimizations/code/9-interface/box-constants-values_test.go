package interfaces

import "testing"

var r interface{}

const N int64 = 12345

func Benchmark_BoxInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = N
	}
}

const F float64 = 1.2345

func Benchmark_BoxFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = F
	}
}

const S = "Go"

func Benchmark_BoxConstString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = S
	}
}
