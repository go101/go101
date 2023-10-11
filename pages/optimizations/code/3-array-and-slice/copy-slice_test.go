package arrays

import "testing"

const N = 64

var r []byte
var s = make([]byte, N)

func init() {
	println("============= N =", N)
	r = make([]byte, N)
}

func copy2(d, s []byte) {
	*(*[N]byte)(d) = *(*[N]byte)(s)
}

func Benchmark_Copy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy(r, s)
	}
}

func Benchmark_Copy2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy2(r, s)
	}
}

func Benchmark_Copy_self(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy(s, s)
	}
}

func Benchmark_Copy2_self(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy2(s, s)
	}
}

func Benchmark_Copy_self_half(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy(s, s[:N/2])
	}
}

