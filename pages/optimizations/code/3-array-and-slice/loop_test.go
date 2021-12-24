package arrays

import (
	"testing"
)

func sum_forrange1(s []int) int {
	var n = 0
	for i := range s {
		n += s[i]
	}
	return n
}

func sum_forrange2(s []int) int {
	var n = 0
	for _, v := range s {
		n += v
	}
	return n
}

//go:noinline
func sum_plainfor(s []int) int {
	var n = 0
	for i := 0; i < len(s); i++ {
		n += s[i]
	}
	return n
}

var s = make([]int, 1<<16)
var r int

func Benchmark_forrange1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = sum_forrange1(s)
	}
}

func Benchmark_forrange2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = sum_forrange2(s)
	}
}

func Benchmark_plainfor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = sum_plainfor(s)
	}
}
