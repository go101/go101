package arrays

import "testing"

var s1 = make([]int, 1<<12)
var s2 = make([]int,len(s1)+1)
var r1 []int
var r2 []int

func Clone_MakeCopy(s []int) []int {
	r := make([]int, len(s))
	copy(r, s)
	return r
}

func Clone_MakeAppend(s []int) []int {
	return append(make([]int, 0, len(s)), s...)
}

func Clone_Append(s []int) []int {
	return append([]int(nil), s...)
}

func Benchmark_Clone_MakeCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r1 = Clone_MakeCopy(s1)
		r2 = Clone_MakeCopy(s2)
	}
}

func Benchmark_Clone_MakeAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r1 = Clone_MakeAppend(s1)
		r2 = Clone_MakeAppend(s2)
	}
}

func Benchmark_Clone_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r1 = Clone_Append(s1)
		r2 = Clone_Append(s2)
	}
}
