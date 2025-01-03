package structs

import "testing"

type T struct {
	A, B, C, D, E, F, G int64
}

//go:noinline
func Zero_ClearIndivials(t *T) {
	t.A, t.C, t.E, t.G = 0, 0, 0, 0
}

//go:noinline
func Zero_Overwrite(t *T) {
	*t = T{B: t.B, D: t.D, F: t.F}
}

var x T

func Benchmark_ClearIndivials(b *testing.B) {
	var y T
	for i := 0; i < b.N; i++ {
		Zero_ClearIndivials(&y)
		x = y
	}
}

func Benchmark_Overwrite(b *testing.B) {
	var y T
	for i := 0; i < b.N; i++ {
		Zero_ClearIndivials(&y)
		x = y
	}
}
