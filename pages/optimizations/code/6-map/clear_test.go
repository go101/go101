package maps

import "testing"

const Max = 1000
var NaN = 0.0

func init() {
	NaN = NaN / NaN
}

func addMapEntries(m map[float64]struct{}, numNaNs int) {
	var k = 1.0
	for i := range [Max]struct{}{} {
		if i < numNaNs {
			m[NaN] = struct{}{}
		} else {
			m[k] = struct{}{}
			k += 1.0
		}
	}
}

func Benchmark_LoopDelete(b *testing.B) {
	var m = make(map[float64]struct{}, Max)
	for i := 0; i < b.N; i++ {
		addMapEntries(m, 0)
		for k := range m {
			delete(m, k)
		}
	}
}

func Benchmark_Clear(b *testing.B) {
	var m = make(map[float64]struct{}, Max)
	for i := 0; i < b.N; i++ {
		addMapEntries(m, 0)
		clear(m)
	}
}

func Benchmark_LoopDelete_WithNaNs(b *testing.B) {
	var m = make(map[float64]struct{}, Max)
	for i := 0; i < b.N; i++ {
		addMapEntries(m, 3)
		for k := range m {
			delete(m, k)
		}
	}
}

func Benchmark_Clear_WithNaNs(b *testing.B) {
	var m = make(map[float64]struct{}, Max)
	for i := 0; i < b.N; i++ {
		addMapEntries(m, 3)
		clear(m)
	}
}

