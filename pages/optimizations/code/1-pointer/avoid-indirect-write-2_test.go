// avoid-indirects_test.go
package pointers

import "testing"

func f(sum *int, s []int) {
	for _, v := range s {
		*sum += v
	}
}

func g(sum *int, s []int) {
	var n = 0
	for _, v := range s {
		n += v
	}
	*sum = n
}

var s = make([]int, 1024)
var r int

func Benchmark_f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(&r, s)
	}
}

func Benchmark_g(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g(&r, s)
	}
}

func h(s []int) int {
	var n = 0
	for _, v := range s {
		n += v
	}
	return n
}

func Benchmark_h(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = h(s)
	}
}
