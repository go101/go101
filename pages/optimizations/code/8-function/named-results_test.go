package functions

import "testing"

const N = 1 << 12

var buf = make([]byte, N)
var r [128][N]byte

func Benchmark_ConvertToArray_Named(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r[i&127] = ConvertToArray_Named(buf)
	}
}

func Benchmark_ConvertToArray_Unnamed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r[i&127] = ConvertToArray_Unnamed(buf)
	}
}

func ConvertToArray_Named(b []byte) (ret [N]byte) {
	// if b == nil {defer print() }
	ret = *(*[N]byte)(b)
	return
}

func ConvertToArray_Unnamed(b []byte) [N]byte {
	// if b == nil {defer print() }
	return *(*[N]byte)(b)
}

func Benchmark_CopyToArray_Named(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r[i&127] = CopyToArray_Named(buf)
	}
}

func Benchmark_CopyToArray_Unnamed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r[i&127] = CopyToArray_Unnamed(buf)
	}
}

func CopyToArray_Named(b []byte) (ret [N]byte) {
	// if b == nil {defer print() }
	copy(ret[:], b)
	return
}

func CopyToArray_Unnamed(b []byte) [N]byte {
	// if b == nil {defer print() }
	var ret [N]byte
	copy(ret[:], b)
	return ret
}
