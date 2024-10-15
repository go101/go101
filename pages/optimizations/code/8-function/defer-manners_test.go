package functions

import "testing"

func foo(x *int) {
	*x++
}

//go:noinline
func f(a int) (r int) {
	defer foo(&r)
	return a
}

//go:noinline
func g(a int) (r int) {
	defer func() { foo(&r) }()
	return a
}

func Benchmark_f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(i)
	}
}

func Benchmark_g(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g(i)
	}
}

//go:noinline
func h(a int) (r int) {
	foo(&a)
	return a
}

func Benchmark_h(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h(i)
	}
}
