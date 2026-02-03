package structs

import "testing"

const N = 1000

type T struct {
	x int
}

//go:noinline
func f(s []int, t *T) {
	for i := range s {
		s[i] = t.x
	}
}

//go:noinline
func g(s []int, t *T) {
	x := t.x
	for i := range s {
		s[i] = x
	}
}

var s = make([]int, 1000)
var t = &T{x: 123}

func Benchmark_f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(s, t)
	}
}

func Benchmark_g(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g(s, t)
	}
}
