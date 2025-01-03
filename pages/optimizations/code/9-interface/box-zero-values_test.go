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

var c64 complex64

func Benchmark_BoxComplex64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = c64
	}
}

var c128 complex128

func Benchmark_BoxComplex128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = c128
	}
}

var str = ""

func Benchmark_BoxString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = str
	}
}

var slc []int

func Benchmark_BoxSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = slc
	}
}

var i interface{}

func Benchmark_BoxInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = i
	}
}

var a [1]struct{x int}

func Benchmark_BoxArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = a
	}
}

var v struct{x [1][]int}

func Benchmark_BoxStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = v
	}
}
