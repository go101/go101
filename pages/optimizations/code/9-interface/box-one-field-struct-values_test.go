package interfaces

import "testing"

var r interface{}

var n8 = struct{ n byte }{0}

func Benchmark_ByteField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = n8
	}
}

var n32 = struct{ n int32 }{255}

func Benchmark_Int32Field(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = n32
	}
}

var f64 = struct{ n float64 }{0}

func Benchmark_Float64Field(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = f64
	}
}

var str = ""

func Benchmark_StringField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = str
	}
}
