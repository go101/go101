package pointers

import "testing"

type T struct {
	n int
	s []int
}

//go:noinline
func (pt *T) mf() {
	pt.n = 0
	for _, v := range pt.s {
		pt.n += v // <=> (*pt).n += v
	}
}

//go:noinline
func (t *T) mg() {
	x := 0
	for _, v := range t.s {
		x += v
	}
	t.n = x
}

const N = 1000

func Benchmark_f(b *testing.B) {
	var t = T{s: make([]int, N)}
	for i := 0; i < b.N; i++ {
		t.mf()
	}
}

func Benchmark_g(b *testing.B) {
	var t = T{s: make([]int, N)}
	for i := 0; i < b.N; i++ {
		t.mg()
	}
}
