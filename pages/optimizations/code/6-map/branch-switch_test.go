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

func MapSwitch(x bool) func() {
	return m[x]
}

func Benchmark_IfElse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IfElse(true)()
		IfElse(false)()
	}
}

func Benchmark_MapSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MapSwitch(true)()
		MapSwitch(false)()
	}
}

func b2i(b bool) (r int) {
	if b {
		r = 1
	}
	return
}

var a = [2]func(){g, f}

func IndexTable(x bool) func() {
	return a[b2i(x)]
}

func Benchmark_IndexTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexTable(true)()
		IndexTable(false)()
	}
}
