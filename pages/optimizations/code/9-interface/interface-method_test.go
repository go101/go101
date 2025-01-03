package interfaces

import "testing"

type BinaryOp interface {
	Do(x, y float64) float64
}

type Add struct{}

func (a Add) Do(x, y float64) float64 {
	return x + y
}

//go:noinline
func (a Add) Do_NotInlined(x, y float64) float64 {
	return x + y
}

var x1, y1, r1 float64
var add Add

func Benchmark_Add_Inline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r1 = add.Do(x1, y1)
	}
}

var x2, y2, r2 float64
var add_NotInlined Add

func Benchmark_Add_NotInlined(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r2 = add.Do_NotInlined(x2, y2)
	}
}

var x3, y3, r3 float64
var add_Interface BinaryOp = Add{}

func Benchmark_Add_Interface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r3 = add_Interface.Do(x3, y3)
	}
}
