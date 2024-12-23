package arrays

import "testing"

const N = 1<<12
var s = make([]int, N)

func Benchmark_copy(b *testing.B) {
	for b.Loop() {
		copy(s, s)
	}
}

func Benchmark_append(b *testing.B) {
	for b.Loop() {
		_ = append(s[:0], s...)
	}
}
