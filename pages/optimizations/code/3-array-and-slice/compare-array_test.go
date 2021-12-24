package arrays

import "testing"

type T [1000]byte

var zero = T{}

func CompareWithLiteral(t *T) bool {
	return *t == T{}
}

func CompareWithGlobalVar(t *T) bool {
	return *t == zero
}

var x T
var r bool

func Benchmark_CompareWithLiteral(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = CompareWithLiteral(&x)
	}
}

func Benchmark_CompareWithGlobalVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = CompareWithGlobalVar(&x)
	}
}
