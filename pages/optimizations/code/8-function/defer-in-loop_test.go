package functions

import "testing"

var n int

func inc() {
	n++
}

func f(n int) {
	for i := 0; i < n; i++ {
		defer inc()
		inc()
	}
}

func g(n int) {
	for i := 0; i < n; i++ {
		func() {
			defer inc()
			inc()
		}()
	}
}

func Benchmark_f(b *testing.B) {
	n = 0
	for i := 0; i < b.N; i++ {
		f(1000)
	}
}

func Benchmark_g(b *testing.B) {
	n = 0
	for i := 0; i < b.N; i++ {
		g(1000)
	}
}
