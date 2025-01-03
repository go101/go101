package copycost

import "testing"

const N = 1024

func buildArray() [N]int {
	var a [N]int
	for i := 0; i < N; i++ {
		a[i] = (N - i) & i
	}
	return a
}

//go:noinline
func Sum_RangeArray(a [N]int) (r int) {
	for _, v := range a {
		r += v
	}
	return
}

//go:noinline
func Sum_RangeArrayPtr1(a *[N]int) (r int) {
	for _, v := range *a {
		r += v
	}
	return
}

//go:noinline
func Sum_RangeArrayPtr2(a *[N]int) (r int) {
	for _, v := range a {
		r += v
	}
	return
}

//go:noinline
func Sum_RangeSlice(a []int) (r int) {
	for _, v := range a {
		r += v
	}
	return
}

func Benchmark_Sum_RangeArray(b *testing.B) {
	var a = buildArray()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sum_RangeArray(a)
	}
}

func Benchmark_Sum_RangeArrayPtr1(b *testing.B) {
	var a = buildArray()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sum_RangeArrayPtr1(&a)
	}
}

func Benchmark_Sum_RangeArrayPtr2(b *testing.B) {
	var a = buildArray()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sum_RangeArrayPtr2(&a)
	}
}

func Benchmark_Sum_RangeSlice(b *testing.B) {
	var a = buildArray()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sum_RangeSlice(a[:])
	}
}
