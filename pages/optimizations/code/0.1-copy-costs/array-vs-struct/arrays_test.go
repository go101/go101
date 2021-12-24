package copycost

import "testing"

var array1_0 [N][1]Element

func Benchmark_CopyArray_1_element(b *testing.B) {
	var array1_1 [N][1]Element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range array1_0 {
			array1_0[k] = array1_1[k]
		}
	}
}

var array2_0 [N][2]Element

func Benchmark_CopyArray_2_elements(b *testing.B) {
	var array2_1 [N][2]Element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range array2_0 {
			array2_0[k] = array2_1[k]
		}
	}
}

var array3_0 [N][3]Element

func Benchmark_CopyArray_3_elements(b *testing.B) {
	var array3_1 [N][3]Element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range array3_0 {
			array3_0[k] = array3_1[k]
		}
	}
}

var array4_0 [N][4]Element

func Benchmark_CopyArray_4_elements(b *testing.B) {
	var array4_1 [N][4]Element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range array4_0 {
			array4_0[k] = array4_1[k]
		}
	}
}

var array5_0 [N][5]Element

func Benchmark_CopyArray_5_elements(b *testing.B) {
	var array5_1 [N][5]Element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range array5_0 {
			array5_0[k] = array5_1[k]
		}
	}
}

var array6_0 [N][6]Element

func Benchmark_CopyArray_6_elements(b *testing.B) {
	var array6_1 [N][6]Element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range array6_0 {
			array6_0[k] = array6_1[k]
		}
	}
}

var array7_0 [N][7]Element

func Benchmark_CopyArray_7_elements(b *testing.B) {
	var array7_1 [N][7]Element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range array7_0 {
			array7_0[k] = array7_1[k]
		}
	}
}

var array8_0 [N][8]Element

func Benchmark_CopyArray_8_elements(b *testing.B) {
	var array8_1 [N][8]Element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range array8_0 {
			array8_0[k] = array8_1[k]
		}
	}
}

var array9_0 [N][9]Element

func Benchmark_CopyArray_9_elements(b *testing.B) {
	var array9_1 [N][9]Element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range array9_0 {
			array9_0[k] = array9_1[k]
		}
	}
}

var array10_0 [N][10]Element

func Benchmark_CopyArray_10_elements(b *testing.B) {
	var array10_1 [N][10]Element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range array10_0 {
			array10_0[k] = array10_1[k]
		}
	}
}

var array11_0 [N][11]Element

func Benchmark_CopyArray_11_elements(b *testing.B) {
	var array11_1 [N][11]Element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range array11_0 {
			array11_0[k] = array11_1[k]
		}
	}
}

var array12_0 [N][12]Element

func Benchmark_CopyArray_12_elements(b *testing.B) {
	var array12_1 [N][12]Element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range array12_0 {
			array12_0[k] = array12_1[k]
		}
	}
}

var array13_0 [N][13]Element

func Benchmark_CopyArray_13_elements(b *testing.B) {
	var array13_1 [N][13]Element
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range array13_0 {
			array13_0[k] = array13_1[k]
		}
	}
}
