package pointers

import "testing"

const N = 1000

var a [N]int
var r int

//go:noinline
func h0(a *[N]int) int {
	var r int
	for i := range a {
		r += a[i]
	}
	return r
}

//go:noinline
func h1(a *[N]int) int {
	var r int
	_ = *a
	for i := range a {
		r += a[i]
	}
	return r
}

//go:noinline
func h2(a *[N]int) int {
	var r int
	s := a[:]
	for i := range s {
		r += s[i]
	}
	return r
}

func Benchmark_h0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = h0(&a)
	}
}

func Benchmark_h1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = h1(&a)
	}
}

func Benchmark_h2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = h2(&a)
	}
}
