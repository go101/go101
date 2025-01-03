package copycost

import "testing"

const N = 1024

type Element = uint64

var a9r [N][9]Element

func Benchmark_CopyArray_9_elements(b *testing.B) {
	var a9s [N][9]Element
	for i := 0; i < b.N; i++ {
		for k := range a9s {
			a9r[k] = a9s[k]
		}
	}
}

var a10r [N][10]Element

func Benchmark_CopyArray_10_elements(b *testing.B) {
	var a10s [N][10]Element
	for i := 0; i < b.N; i++ {
		for k := range a10s {
			a10r[k] = a10s[k]
		}
	}
}

type S9 struct{ a, b, c, d, e, f, g, h, i Element }

var s9r [N]S9

func Benchmark_CopyStruct_9_fields(b *testing.B) {
	var s9s [N]S9
	for i := 0; i < b.N; i++ {
		for k := range s9s {
			s9r[k] = s9s[k]
		}
	}
}

type S10 struct{ a, b, c, d, e, f, g, h, i, j Element }

var s10r [N]S10

func Benchmark_CopyStruct_10_fields(b *testing.B) {
	var s10s [N]S10
	for i := 0; i < b.N; i++ {
		for k := range s10s {
			s10r[k] = s10s[k]
		}
	}
}
