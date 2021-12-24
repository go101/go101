package copycost

import "testing"

var e0 [N]Element

func Benchmark_Scalar(b *testing.B) {
	var e1 Element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range e0 {
			e0[k] = e1
		}
	}
}
