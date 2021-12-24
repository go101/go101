package structs

import "testing"

const N = 1000

type T struct {
	x int
}

//go:noinline
func f(t *T) {
	t.x = 0
	for i := 0; i < N; i++ {
		t.x += i
	}
}

//go:noinline
func g(t *T) {
	var x = 0
	for i := 0; i < N; i++ {
		x += i
	}
	t.x = x
}

var t = &T{}

func Benchmark_f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(t)
	}
}

func Benchmark_g(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g(t)
	}
}

//go:noinline
func h(t *T) {
	x := &t.x
	for i := 0; i < N; i++ {
		*x += i
	}
}

func Benchmark_h(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h(t)
	}
}
