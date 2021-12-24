package interfaces

import "testing"

var r interface{}

func Benchmark_BoxSmallInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = int16(i & 255)
	}
}

func Benchmark_BoxSmallInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = int32(i & 255)
	}
}

func Benchmark_BoxSmallInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = int32(i & 255)
	}
}
