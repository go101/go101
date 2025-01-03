package bce

import "testing"

var r byte
var s = make([]byte, 1000)
const N = 3

func f6a(s []byte) {
	for i := 0; i < len(s)-(N-1); i += N {
		r = s[i+N-1] // Found IsInBounds
	}
}

func f6b(s []byte) {
	for i := N-1; i < len(s); i += N {
		r = s[i] // Found IsInBounds
	}
}

func f6d(s []byte) {
	var k = uint(len(s)-(N-1))
	for i := uint(0); i < k; i += N {
		r = s[i+N-1] // Found IsInBounds
	}
}

func f6c(s []byte) {
	for i := uint(N-1); i < uint(len(s)); i += N {
		r = s[i]
	}
}



func Benchmark_f6a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f6a(s)
	}
}

func Benchmark_f6b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f6b(s)
	}
}

func Benchmark_f6c(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f6c(s)
	}
}

func Benchmark_f6d(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f6d(s)
	}
}



