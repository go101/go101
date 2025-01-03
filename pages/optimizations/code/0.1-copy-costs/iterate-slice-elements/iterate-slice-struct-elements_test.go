package copycost

import "testing"

type S struct{ a, b, c, d, e int }

//go:noinline
func sum_UseSecondIterationVar(s []S) int {
	var sum int
	for _, v := range s {
		sum += v.c
		sum += v.d
		sum += v.e
	}
	return sum
}

//go:noinline
func sum_OneIterationVar_Index(s []S) int {
	var sum int
	for i := range s {
		sum += s[i].c
		sum += s[i].d
		sum += s[i].e
	}
	return sum
}

//go:noinline
func sum_OneIterationVar_Ptr(s []S) int {
	var sum int
	for i := range s {
		v := &s[i]
		sum += v.c
		sum += v.d
		sum += v.e
	}
	return sum
}

var s = make([]S, 1000)
var r [128]int

func init() {
	for i := range s {
		s[i] = S{i, i, i, i, i}
	}
}

func Benchmark_UseSecondIterationVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r[i&127] = sum_UseSecondIterationVar(s)
	}
}

func Benchmark_OneIterationVar_Index(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r[i&127] = sum_OneIterationVar_Index(s)
	}
}

func Benchmark_OneIterationVar_Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r[i&127] = sum_OneIterationVar_Ptr(s)
	}
}
