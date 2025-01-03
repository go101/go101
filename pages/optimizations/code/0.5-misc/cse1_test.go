package functions

import "testing"

type T [1 << 13]byte

var r, t [1]T

const N = 0

var n = 0

func Benchmark_InvertBits_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := &r[n]
		t := &t[n]
		for k := 0; k < len(T{}); k++ {
			r[k] = t[k]
		}
	}
}

func Benchmark_InvertBits_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := &r[N]
		t := &t[N]
		for k := 0; k < len(T{}); k++ {
			r[k] = t[k]
		}
	}
}

var r0, t0 T

func Benchmark_InvertBits_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := 0; k < len(T{}); k++ {
			r0[k] = t0[k]
		}
	}
}
