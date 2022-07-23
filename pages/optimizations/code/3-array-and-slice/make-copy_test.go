package arrays

import "testing"

// Before Go toolchain 1.19, the length used here is 2048.
// But sometiems, the length 2048 is unable to verify the assuptions
// with Go toolchain 1.19.
// No sure whether or not this is caused by the runtime changes added in 1.19.
// Todo: need investigations.
var s = make([]int, 2048)
var a = [1][]int{s}
var t = struct{x []int}{s}
var r []int

func Benchmark________________(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = make([]int, len(s))
		copy(r, s)
	}
}

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
	f := func(interface{}) {}
	for i := 0; i < b.N; i++ {
		r = make([]int, len(s))
		f(copy(r, s))
	}
}

func Benchmark_MakeCopy_f(b *testing.B) {
	var k interface{}
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

func Benchmark_MakeCopy4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var a, b = s[:0], s[0:]
		r = make([]int, len(s))
		copy(r, a)
		copy(r[len(a):], b)
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
