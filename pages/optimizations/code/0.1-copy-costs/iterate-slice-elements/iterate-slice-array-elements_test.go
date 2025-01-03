package copycost

import "testing"

type Element [10]int64

//go:noinline
func Sum_PlainForLoop(s []Element) (r int64) {
	for i := 0; i < len(s); i++ {
		r += s[i][0]
	}
	return
}

//go:noinline
func Sum_OneIterationVar(s []Element) (r int64) {
	for i := range s {
		r += s[i][0]
	}
	return
}

//go:noinline
func Sum_UseSecondIterationVar(s []Element) (r int64) {
	for _, v := range s {
		r += v[0]
	}
	return
}

//===================

func buildSlice() []Element {
	var s = make([]Element, 1000)
	for i := 0; i < len(s); i++ {
		s[i] = Element{0: int64(i)}
	}
	return s
}

var r [128]int64

func Benchmark_PlainForLoop(b *testing.B) {
	var s = buildSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r[i&127] = Sum_PlainForLoop(s)
	}
}

func Benchmark_OneIterationVar(b *testing.B) {
	var s = buildSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r[i&127] = Sum_OneIterationVar(s)
	}
}

func Benchmark_UseSecondIterationVar(b *testing.B) {
	var s = buildSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r[i&127] = Sum_UseSecondIterationVar(s)
	}
}
