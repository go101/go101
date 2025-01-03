package functions

import "testing"

type T [100]byte

var gt T

//go:noinline
func f_noninline() T {
	return gt
}

func f_autoinline() T {
	return gt
}

func Benchmark_f_noninline(b *testing.B) {
	var t T
	for i := 0; i < b.N; i++ {
		t = f_noninline()
	}
	gt = t
}

func Benchmark_f_autoinline(b *testing.B) {
	var t T
	for i := 0; i < b.N; i++ {
		t = f_autoinline()
	}
	gt = t
}

func Benchmark_f_manualinline(b *testing.B) {
	var t T
	for i := 0; i < b.N; i++ {
		t = gt
	}
	gt = t
}

//go:noinline
func g_noinline() T {
	return T{}
}

func g_autoinline() T {
	return T{}
}

func Benchmark_g_noinline(b *testing.B) {
	var t T
	for i := 0; i < b.N; i++ {
		t = g_noinline()
	}
	gt = t
}

func Benchmark_g_autoinline(b *testing.B) {
	var t T
	for i := 0; i < b.N; i++ {
		t = g_autoinline()
	}
	gt = t
}

func Benchmark_g_manualinline(b *testing.B) {
	var t T
	for i := 0; i < b.N; i++ {
		t = T{}
	}
	gt = t
}

//go:noinline
func h_noinline(t *T) {
	*t = gt
}

func h_autoinline(t *T) {
	*t = gt
}

func Benchmark_h_noinline(b *testing.B) {
	var t T
	for i := 0; i < b.N; i++ {
		h_noinline(&t)
	}
	gt = t
}

func Benchmark_h_autoinline(b *testing.B) {
	var t T
	for i := 0; i < b.N; i++ {
		h_autoinline(&t)
	}
	gt = t
}

func Benchmark_h_manualinline(b *testing.B) {
	var t T
	var pt = &t
	for i := 0; i < b.N; i++ {
		*pt = gt
	}
	gt = t
}
