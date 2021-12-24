package interfaces

import "testing"

var r interface{}

var p = new([100]int)

func Benchmark_BoxPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = p
	}
}

var m = map[string]int{"Go": 2009}

func Benchmark_BoxMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = m
	}
}

var c = make(chan int, 100)

func Benchmark_BoxChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = c
	}
}

var f = func(a, b int) int { return a + b }

func Benchmark_BoxFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = f
	}
}
