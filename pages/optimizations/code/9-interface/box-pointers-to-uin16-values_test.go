package interfaces

import "testing"

var values [65536]uint16

func init() {
	for i := range values {
		values[i] = uint16(i)
	}
}

var r interface{}

func Benchmark_Box_Normal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = uint16(i)
	}
}

func Benchmark_Box_Lookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = &values[uint16(i)]
	}
}

var x interface{} = uint16(10000)
var y interface{} = &values[10000]
var n uint16

func Benchmark_Unbox_Normal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n = x.(uint16)
	}
}

func Benchmark_Unbox_Lookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n = *y.(*uint16)
	}
}
