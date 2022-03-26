package bce

import (
	"testing"
)

var n = 1 << 11
var r []int

func f9a(n int) []int {
	buf := make([]int, n+1)
	k := 0
	for i := 0; i <= n; i++ {
		buf[i] = k // Found IsInBounds
		k++
	}
	return buf
}

func f9b(n int) []int {
	buf := make([]int, n+1)
	k := 0
	for i := 0; i < len(buf); i++ {
		buf[i] = k
		k++
	}
	return buf
}

func Benchmark_f9a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = f9a(n)
	}
}

func Benchmark_f9b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = f9b(n)
	}
}
