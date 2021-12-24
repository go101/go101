package interfaces

import "testing"

var r interface{}

var n16 int16 = 12345

func Benchmark_BoxInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = n16
	}
}

var n32 int32 = 12345

func Benchmark_BoxInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = n32
	}
}

var n64 int64 = 12345

func Benchmark_BoxInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = n64
	}
}

var f64 float64 = 1.2345

func Benchmark_BoxFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = f64
	}
}

var s = "Go"

func Benchmark_BoxString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = s
	}
}

var x = []int{1, 2, 3}

func Benchmark_BoxSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = x
	}
}

var a = [100]int{}

func Benchmark_BoxArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = a
	}
}
