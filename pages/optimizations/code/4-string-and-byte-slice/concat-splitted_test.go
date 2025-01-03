package bytes

import "testing"

var s37 = []byte{36: 'x'} // len(s37) == 37
var str string

func Benchmark_concat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = string(s37) + string(s37)
	}
}

func Benchmark_concat_splited(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = string(s37[:32]) +
			string(s37[32:]) +
			string(s37[:32]) +
			string(s37[32:])
	}
}
