package bytes

import "testing"

var s = "1234567890abcdef" // len(s) == 16
var r string

func Concat_WithPlus(a, b, c, d string) string {
	return a + b + c + d
}

func Concat_WithBytes(ss ...string) string {
	var n = 0
	for _, s := range ss {
		n += len(s)
	}
	var bs []byte
	if n > 64 {
		bs = make([]byte, 0, n) // escapes to heap
	} else {
		bs = make([]byte, 0, 64) // does not escape
	}
	for _, s := range ss {
		bs = append(bs, s...)
	}
	return string(bs)
}

func Benchmark_Concat_WithPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Concat_WithPlus(s, s, s, s)
	}
}

func Benchmark_Concat_WithBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Concat_WithBytes(s, s, s, s)
	}
}
