package functions

import "testing"

var sum int

func f(s []int) {
	for _, v := range s {
		sum += v
	}
}

func g(s []int) {
	var n = 0
	for _, v := range s {
		n += v
	}
	sum = n
}

var s = make([]int, 1024)

func Benchmark_f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(s)
	}
}
func Benchmark_g(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g(s)
	}
}
