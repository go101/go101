// unnecessary-checks.go
package pointers

import "testing"

const N = 1000

var a [N]int

//go:noinline
func g1(a *[N]int) {
	_ = *a
	for i := range a {
		a[i] = i
	}
}

//go:noinline
func g0(a *[N]int) {
	for i := range a {
		a[i] = i
	}
}

func Benchmark_g0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g0(&a)
	}
}

func Benchmark_g1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g1(&a)
	}
}

//go:noinline
func g2(x *[N]int) {
	a := x[:]
	for i := range a {
		a[i] = i
	}
}

func Benchmark_g2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g2(&a)
	}
}
