package arrays

import "testing"

// https://github.com/golang/go/issues/53888

func Grow_MakeCopy(x []byte, c int) []byte {
	r := make([]byte, c)
	copy(r, x)
	return r[:len(x)]
}

func Grow_MakeCopy2(x []byte, c int) []byte {
	r := make([]byte, c, c)
	_ = copy(r, x)
	return r[:len(x)]
}

func Grow_Oneline(x []byte, c int) []byte {
	return append(x, make([]byte, c - len(x))...)[:len(x)]
}

func Grow_Oneline_Named(x []byte, c int) []byte {
	type T []byte
	return append(x, make(T, c - len(x))...)[:len(x)]
}

var n = 1000
var m = n + 1000

// go official benchmark tool is unable to get the expected results
// when the element type is byte.
type byte = int

var s = make([]byte, n)
var r []byte

func Benchmark_______warmup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Grow_MakeCopy(s, m)
	}
}

func Benchmark_Grow_MakeCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Grow_MakeCopy(s, m)
	}
}

func Benchmark_Grow_MakeCopy2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Grow_MakeCopy2(s, m)
	}
}

func Benchmark_Grow_Oneline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Grow_Oneline(s, m)
	}
}

func Benchmark_Grow_Oneline_Named(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Grow_Oneline_Named(s, m)
	}
}

