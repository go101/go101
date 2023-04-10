package maps

import "testing"

//go:noiline
func f() {}

//go:noiline
func g() {}

func IfElse(x bool) func() {
	if x {
		return f
	} else {
		return g
	}
}

var m = map[bool]func(){true: f, false: g}

func Benchmark_IfElse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IfElse(true)()
		IfElse(false)()
	}
}

func Benchmark_MapSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m[true]()
		m[false]()
	}
}

var a = [2]func(){g, f}

func IndexTable(x bool) func() {
	if x {
		return a[1]
	}
	return a[0]
}

func Benchmark_IndexTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexTable(true)()
		IndexTable(false)()
	}
}

func b2i(b bool) (r int) {
	if b {
		r = 1
	}
	return
}

var boolMap = [2]func(){g, f}

func Benchmark_BoolMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		boolMap[b2i(true)]()
		boolMap[b2i(false)]()
	}
}

func Benchmark_BoolMap2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		boolMap[b2i(true)&1]()
		boolMap[b2i(false)&1]()
	}
}
