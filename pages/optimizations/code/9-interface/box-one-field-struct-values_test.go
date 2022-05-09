package interfaces

import "testing"

var r interface{}

var zs = struct{x struct{}}{}

func Benchmark_ZeroSizeFieldField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = zs
	}
}

var n8 = struct{ n byte }{0}

func Benchmark_ByteField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = n8
	}
}

func Benchmark_BoolField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = struct{ n bool }{}
	}
}

func Benchmark_SmallInt32Field(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = struct{ n int32 }{255}
	}
}

func Benchmark_Int32Field(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = struct{ n int32 }{256}
	}
}

func Benchmark_ZeroFloat64Element(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = [1]float64{0}
	}
}

func Benchmark_Float64Element(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = [1]float64{1.23}
	}
}

func Benchmark_BlankStringField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = struct{ n string }{""}
	}
}

func Benchmark_StringField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = struct{ n string }{"abc"}
	}
}

func Benchmark_NilSliceField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = struct{ n []int }{nil}
	}
}

func Benchmark_SliceField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = struct{ n []int }{[]int{}}
	}
}

func Benchmark_PointerField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = struct{n *int}{new(int)}
	}
}

func Benchmark_NilFunctionElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = [1]func(){func(){}}
	}
}

var m = struct{ n map[int]int }{map[int]int{1: 2}}

func Benchmark_MapField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = m
	}
}

func Benchmark_ArrayField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = struct{ n [2]string }{}
	}
}

func Benchmark_StructElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = [1]struct{ n, m string }{}
	}
}
