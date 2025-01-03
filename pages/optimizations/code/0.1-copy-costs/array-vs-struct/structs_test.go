package copycost

import "testing"

var struct1_0 [N]struct{ a Element }

func Benchmark_CopyStruct_1_field(b *testing.B) {
	var struct1_1 struct{ a Element }
	for i := 0; i < b.N; i++ {
		for k := range struct1_0 {
			struct1_0[k] = struct1_1
		}
	}
}

var struct2_0 [N]struct{ a, b Element }

func Benchmark_CopyStruct_2_fields(b *testing.B) {
	var struct2_1 struct{ a, b Element }
	for i := 0; i < b.N; i++ {
		for k := range struct2_0 {
			struct2_0[k] = struct2_1
		}
	}
}

var struct3_0 [N]struct{ a, b, c Element }

func Benchmark_CopyStruct_3_fields(b *testing.B) {
	var struct3_1 struct{ a, b, c Element }
	for i := 0; i < b.N; i++ {
		for k := range struct3_0 {
			struct3_0[k] = struct3_1
		}
	}
}

var struct4_0 [N]struct{ a, b, c, d Element }

func Benchmark_CopyStruct_4_fields(b *testing.B) {
	var struct4_1 struct{ a, b, c, d Element }
	for i := 0; i < b.N; i++ {
		for k := range struct4_0 {
			struct4_0[k] = struct4_1
		}
	}
}

var struct5_0 [N]struct{ a, b, c, d, e Element }

func Benchmark_CopyStruct_5_fields(b *testing.B) {
	var struct5_1 struct{ a, b, c, d, e Element }
	for i := 0; i < b.N; i++ {
		for k := range struct5_0 {
			struct5_0[k] = struct5_1
		}
	}
}

var struct6_0 [N]struct{ a, b, c, d, e, f Element }

func Benchmark_CopyStruct_6_fields(b *testing.B) {
	var struct6_1 struct{ a, b, c, d, e, f Element }
	for i := 0; i < b.N; i++ {
		for k := range struct6_0 {
			struct6_0[k] = struct6_1
		}
	}
}

var struct7_0 [N]struct{ a, b, c, d, e, f, g Element }

func Benchmark_CopyStruct_7_fields(b *testing.B) {
	var struct7_1 struct{ a, b, c, d, e, f, g Element }
	for i := 0; i < b.N; i++ {
		for k := range struct7_0 {
			struct7_0[k] = struct7_1
		}
	}
}

var struct8_0 [N]struct{ a, b, c, d, e, f, g, h Element }

func Benchmark_CopyStruct_8_fields(b *testing.B) {
	var struct8_1 struct{ a, b, c, d, e, f, g, h Element }
	for i := 0; i < b.N; i++ {
		for k := range struct8_0 {
			struct8_0[k] = struct8_1
		}
	}
}

var struct9_0 [N]struct{ a, b, c, d, e, f, g, h, i Element }

func Benchmark_CopyStruct_9_fields(b *testing.B) {
	var struct9_1 struct{ a, b, c, d, e, f, g, h, i Element }
	for i := 0; i < b.N; i++ {
		for k := range struct9_0 {
			struct9_0[k] = struct9_1
		}
	}
}

var struct10_0 [N]struct{ a, b, c, d, e, f, g, h, i, j Element }

func Benchmark_CopyStruct_10_fields(b *testing.B) {
	var struct10_1 struct{ a, b, c, d, e, f, g, h, i, j Element }
	for i := 0; i < b.N; i++ {
		for k := range struct10_0 {
			struct10_0[k] = struct10_1
		}
	}
}
