package arrays

import "testing"

var s = make([]int, 2048)
var a = [1][]int{s}
var t = struct{x []int}{s}
var r []int

func Benchmark_MakeCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = make([]int, len(s))
		copy(r, s)
	}
}

func Benchmark_MakeCopy_b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = make([]int, len(s), len(s))
		copy(r, s)
	}
}

func Benchmark_MakeCopy_c(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = make([]int, len(a[0]))
		copy(r, a[0])
	}
}

func Benchmark_MakeCopy_d(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = make([]int, len(t.x))
		copy(r, t.x)
	}
}

func Benchmark_MakeCopy_e(b *testing.B) {
	f := func(int) {}
	for i := 0; i < b.N; i++ {
		r = make([]int, len(s))
		f(copy(r, s))
	}
}

func Benchmark_MakeCopy_f(b *testing.B) {
	var k int
	for i := 0; i < b.N; i++ {
		r = make([]int, len(s))
		k = copy(r, s)
	}
	_ = k
}

func Benchmark_MakeCopy_g(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = make([]int, len(s))
		_ = copy(r, s)
	}
}

func Benchmark_MakeCopy2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var a, b = s[:len(s)], s[len(s):]
		r = make([]int, len(s))
		copy(r, a)
		copy(r[len(a):], b)
	}
}

func Benchmark_MakeCopy3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var a, b = s[:len(s)/2], s[len(s)/2:]
		r = make([]int, len(s))
		copy(r, a)
		copy(r[len(a):], b)
	}
}

func Benchmark_MakeCopy4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var a, b = s[:0], s[0:]
		r = make([]int, len(s))
		copy(r, a)
		copy(r[len(a):], b)
	}
}
