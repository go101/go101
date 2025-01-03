package arrays

import "testing"

// Before Go toolchain 1.19, the length has no effects on the (relative) benchmark results.
// Since 1.19, we can only get expected benchmarks when the length is out of range of (128, 4000).
// No sure whether or not this is caused by the runtime changes added in 1.19.
// Todo: need investigations.
var n = 4000
var s = make([]int, n)
var a = [1][]int{s}
var t = struct{x []int}{s}
var r []int

func Benchmark________________(b *testing.B) {
	for b.Loop() {
		r = make([]int, len(s))
		copy(r, s)
	}
}

func Benchmark_MakeCopy(b *testing.B) {
	for b.Loop() {
		r = make([]int, len(s))
		copy(r, s)
	}
}

func Benchmark_MakeCopy_b(b *testing.B) {
	for b.Loop() {
		r = make([]int, len(s), len(s))
		copy(r, s)
	}
}

func Benchmark_MakeCopy_c(b *testing.B) {
	for b.Loop() {
		r = make([]int, len(a[0]))
		copy(r, a[0])
	}
}

func Benchmark_MakeCopy_d(b *testing.B) {
	for b.Loop() {
		r = make([]int, len(t.x))
		copy(r, t.x)
	}
}

func Benchmark_MakeCopy_e(b *testing.B) {
	f := func(interface{}) {}
	for b.Loop() {
		r = make([]int, len(s))
		f(copy(r, s))
	}
}

func Benchmark_MakeCopy_f(b *testing.B) {
	var k interface{}
	for b.Loop() {
		r = make([]int, len(s))
		k = copy(r, s)
	}
	_ = k
}

func Benchmark_MakeCopy_g(b *testing.B) {
	for b.Loop() {
		r = make([]int, len(s))
		_ = copy(r, s)
	}
}

var x, y []int

func Benchmark_MakeCopy2(b *testing.B) {
	for b.Loop() {
		x, y = s[:len(s)], s[len(s):]
		r = make([]int, len(x))
		copy(r, x)
		copy(r[len(x):], y)
	}
}

func Benchmark_MakeCopy3(b *testing.B) {
	for b.Loop() {
		x, y = s[:len(s)/2], s[len(s)/2:]
		r = make([]int, len(s))
		copy(r, x)
		copy(r[len(x):], y)
	}
}

func Benchmark_MakeCopy4(b *testing.B) {
	for b.Loop() {
		x, y = s[:0], s[0:]
		r = make([]int, len(s))
		copy(r, x)
		copy(r[len(x):], y)
	}
}
