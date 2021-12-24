package interfaces

import "testing"

var r interface{}

var f32 float32 = 0

func Benchmark_BoxFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = f32
	}
}

var f64 float64 = 0

func Benchmark_BoxFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = f64
	}
}

var s = ""

func Benchmark_BoxString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = s
	}
}
