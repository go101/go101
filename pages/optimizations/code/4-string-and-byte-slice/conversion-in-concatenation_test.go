package bytes

import "testing"

var bs = []byte{62: 'x'} // len(bs) == 37
var str string

func Benchmark_concat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = string(bs) + string(bs)
	}
}

func Benchmark_concat_split(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = string(bs[:32]) +
			string(bs[32:]) +
			string(bs[:32]) +
			string(bs[32:])
	}
}

func Benchmark_concat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = (" " + string(bs) + string(bs))[1:]
	}
}

func Benchmark_concat2_strings(b *testing.B) {
	var s = string(bs)
	for i := 0; i < b.N; i++ {
		str = (" " + s + s)[1:]
	}
}
