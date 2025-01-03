package functions

import "testing"

type T [1 << 8]byte

var r, s T

//go:noinline
func not_inline_able(x1, y1 *T) {
	x, y := x1[:], y1[:]
	for k := 0; k < len(T{}); k++ {
		x[k] = y[k]
	}
}

func inline_able(x1, y1 *T) {
	x, y := x1[:], y1[:]
	for k := 0; k < len(T{}); k++ {
		x[k] = y[k]
	}
}

func Benchmark_not_inlined(b *testing.B) {
	for i := 0; i < b.N; i++ {
		not_inline_able(&r, &s)
	}
}

func Benchmark_auto_inlined(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inline_able(&r, &s)
	}
}

func Benchmark_manual_inlined(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := 0; k < len(T{}); k++ {
			r[k] = s[k]
		}
	}
}
