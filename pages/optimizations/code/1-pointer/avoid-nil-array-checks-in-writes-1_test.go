package pointers

import "testing"

const N = 1000

type T struct {
	a *[N]int
}

//go:noinline
func f0(t *T) {
	for i := range t.a {
		t.a[i] = i
	}
}

//go:noinline
func f1(t *T) {
	_ = *t.a
	for i := range t.a {
		t.a[i] = i
	}
}

//go:noinline
func f2(t *T) {
	a := t.a
	for i := range a {
		a[i] = i
	}
}

//go:noinline
func f3(t *T) {
	a := t.a
	_ = *a
	for i := range a {
		a[i] = i
	}
}

//go:noinline
func f4(t *T) {
	a := t.a[:]
	for i := range a {
		a[i] = i
	}
}

func Benchmark_f0(b *testing.B) {
	var t = &T{a: new([N]int)}
	for i := 0; i < b.N; i++ {
		f0(t)
	}
}

func Benchmark_f1(b *testing.B) {
	var t = &T{a: new([N]int)}
	for i := 0; i < b.N; i++ {
		f1(t)
	}
}

func Benchmark_f2(b *testing.B) {
	var t = &T{a: new([N]int)}
	for i := 0; i < b.N; i++ {
		f2(t)
	}
}

func Benchmark_f3(b *testing.B) {
	var t = &T{a: new([N]int)}
	for i := 0; i < b.N; i++ {
		f3(t)
	}
}

func Benchmark_f4(b *testing.B) {
	var t = &T{a: new([N]int)}
	for i := 0; i < b.N; i++ {
		f4(t)
	}
}
