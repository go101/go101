package pointers

import "testing"

func f(v *int, s []int) {
	for i := range s {
		s[i] = *v
	}
}

func g(v *int, s []int) {
	var n = *v
	for i := range s {
		s[i] = n
	}
}

var s = make([]int, 1024)
var r int = 123

func Benchmark_f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(&r, s)
	}
}

func Benchmark_g(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g(&r, s)
	}
}
