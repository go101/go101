// avoid-indirects_test.go
package pointers

import "testing"

//go:noinline
func f(sum *int, s []int) {
	for _, v := range s { // line 8
		*sum += v // line 9
	}
}

//go:noinline
func g(sum *int, s []int) {
	var n = *sum
	for _, v := range s { // line 16
		n += v // line 17
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

//go:noinline
func h(s []int) int {
	var n = 0
	for _, v := range s {
		n += v
	}
	return n
}

func Benchmark_h(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r += h(s)
	}
}

func init() {
	{
		var s = []int{1, 1, 1}
		var sum = &s[2]
		f(sum, s)
		println(*sum) // 6
	}

	{
		var s = []int{1, 1, 1}
		var sum = &s[2]
		g(sum, s)
		println(*sum) // 4
	}

	{
		var s = []int{1, 1, 1}
		var sum = &s[2]
		*sum = h(s)
		println(*sum) // 3
	}
}
