package pointers

import "testing"

const N = 1000

var r int

type T struct {
	a *[N]int
}

//go:noinline
func f0(t *T) int {
	var r int
	for i := range t.a {
		r += t.a[i]
	}
	return r
}

//go:noinline
func f1(t *T) int {
	var r int
	_ = *t.a
	for i := range t.a {
		r += t.a[i]
	}
	return r
}

//go:noinline
func f2(t *T) int {
	var r int
	a := t.a
	for i := range a {
		r += a[i]
	}
	return r
}

//go:noinline
func f3(t *T) int {
	var r int
	a := t.a
	_ = *a
	for i := range a {
		r += a[i]
	}
	return r
}

//go:noinline
func f4(t *T) int {
	var r int
	a := t.a[:]
	for i := range a {
		r += a[i]
	}
	return r
}

func Benchmark_f0(b *testing.B) {
	var t = &T{a: new([N]int)}
	for i := 0; i < b.N; i++ {
		r = f0(t)
	}
}

func Benchmark_f1(b *testing.B) {
	var t = &T{a: new([N]int)}
	for i := 0; i < b.N; i++ {
		r = f1(t)
	}
}

func Benchmark_f2(b *testing.B) {
	var t = &T{a: new([N]int)}
	for i := 0; i < b.N; i++ {
		r = f2(t)
	}
}

func Benchmark_f3(b *testing.B) {
	var t = &T{a: new([N]int)}
	for i := 0; i < b.N; i++ {
		r = f3(t)
	}
}

func Benchmark_f4(b *testing.B) {
	var t = &T{a: new([N]int)}
	for i := 0; i < b.N; i++ {
		r = f4(t)
	}
}
