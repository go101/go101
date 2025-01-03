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

func h(s []int) {
	var n [2]int
	for _, v := range s {
		n[0] += v
	}
	sum = n[0]
}

func p(s []int) {
	var n = new(int)
	for _, v := range s {
		*n += v
	}
	sum = *n
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

func Benchmark_h(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h(s)
	}
}

func Benchmark_p(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p(s)
	}
}
