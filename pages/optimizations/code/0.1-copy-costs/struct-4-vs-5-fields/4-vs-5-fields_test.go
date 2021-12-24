package copycost

import "testing"

type T4 struct{ a, b, c, d float32 }
type T5 struct{ a, b, c, d, e float32 }

var t4 T4
var t5 T5

//go:noinline
func Add4(x, y T4) (z T4) {
	type _ int // avoiding being inlined
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	z.d = x.d + y.d
	return
}

//go:noinline
func Add5(x, y T5) (z T5) {
	type _ int // avoiding being inlined
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	z.d = x.d + y.d
	z.e = x.e + y.e
	return
}

func Benchmark_Add4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x, y T4
		t4 = Add4(x, y)
	}
}

func Benchmark_Add5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x, y T5
		t5 = Add5(x, y)
	}
}
