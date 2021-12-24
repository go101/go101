package structs

import "testing"

const M = 1024

type Element = uint64

var struct2_0 [M]struct{ a, b Element }

func Benchmark_CopyStruct_2_fields(b *testing.B) {
	var struct2_1 [M]struct{ a, b Element }
	for i := 0; i < b.N; i++ {
		for k := range struct2_0 {
			struct2_0[k] = struct2_1[k]
		}
	}
}

var array2_0 [M][2]Element

func Benchmark_CopyArray_2_elements(b *testing.B) {
	var array2_1 [M][2]Element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range array2_0 {
			array2_0[k] = array2_1[k]
		}
	}
}
