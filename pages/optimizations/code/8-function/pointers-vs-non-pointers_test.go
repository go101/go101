package functions

import "testing"

type T5 struct{ a, b, c, d, e float32 }

var t5 T5

//go:noinline
func Add5_TT_T(x, y T5) (z T5) {
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	z.d = x.d + y.d
	z.e = x.e + y.e
	return
}

//go:noinline
func Add5_TT_P(z *T5, x, y T5) {
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	z.d = x.d + y.d
	z.e = x.e + y.e
}

//go:noinline
func Add5_PP_T(x, y *T5) (z T5) {
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	z.d = x.d + y.d
	z.e = x.e + y.e
	return
}

//go:noinline
func Add5_PPP(z, x, y *T5) {
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	z.d = x.d + y.d
	z.e = x.e + y.e
}

func Benchmark_Add5_TT_T(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x, y, z T5
		z = Add5_TT_T(x, y)
		t5 = z
	}
}

func Benchmark_Add5_TT_P(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x, y, z T5
		Add5_TT_P(&z, x, y)
		t5 = z
	}
}

func Benchmark_Add5_PP_T(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x, y, z T5
		z = Add5_PP_T(&x, &y)
		t5 = z
	}
}

func Benchmark_Add5_PPP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x, y, z T5
		Add5_PPP(&z, &x, &y)
		t5 = z
	}
}

type T4 struct{ a, b, c, d float32 }

var t4 T4

//go:noinline
func Add4_TT_T(x, y T4) (z T4) {
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	z.d = x.d + y.d
	return
}

//go:noinline
func Add4_PPP(z, x, y *T4) {
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	z.d = x.d + y.d
}

func Benchmark_Add4_TT_T(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x, y, z T4
		z = Add4_TT_T(x, y)
		t4 = z
	}
}

func Benchmark_Add4_PPP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x, y, z T4
		Add4_PPP(&z, &x, &y)
		t4 = z
	}
}
