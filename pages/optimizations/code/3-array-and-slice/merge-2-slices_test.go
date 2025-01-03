package arrays

import "testing"

const M, N = 1000, 3000
var x = make([]byte, M)
var y = make([]byte, N)
var r []byte

func init() {
	println("===", cap(Merge_MakeCopy(x, y)), cap(Merge_Append(x, y)))
}

func Merge_MakeCopy(a, b []byte) []byte {
	r := make([]byte, len(a) + len(b))
	copy(r, a)
	copy(r[len(a):], b)
	return r
}

func Merge_Append(a, b []byte) []byte {
	return append(a, b...)
}

func Benchmark_Merge_MakeCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Merge_MakeCopy(x, y)
	}
}

func Benchmark_Merge_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Merge_Append(x, y)
	}
}

