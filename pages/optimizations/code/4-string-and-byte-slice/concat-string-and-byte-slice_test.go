package bytes

import "testing"

const N = 8
const M = 1000

var str = string(make([]byte, N))
var bs = make([]byte, M)

var x []byte

func Benchmark_makecopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var newByteSlice = make([]byte, len(str)+len(bs))
		copy(newByteSlice, str)
		copy(newByteSlice[len(str):], bs)
		x = newByteSlice
	}
}

func Benchmark_append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = append([]byte(str), bs...)
	}
}
