package functions

import "testing"

const N = 100

// This function is inline-able.
func Slice2Array(b []byte) [N]byte {
	return *(*[N]byte)(b)
}

//====================

var buf = make([]byte, 8192)
var r [128][N]byte

func Benchmark_ManualInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r[i&127] = *(*[N]byte)(buf)
	}
}

func Benchmark_AutoInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r[i&127] = Slice2Array(buf) // inlined
	}
}
