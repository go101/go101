package allocations

import "testing"
import "strings"

var str1 = strings.Repeat("x", 60)
var str2 = strings.Repeat("x", 6000)
var bs = make([]byte, 6000)

var r1 []byte

func Benchmark_ConvertAppend_SmallString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r1 = append([]byte(str1), bs...)
	}
}

var r2 []byte

func Benchmark_MakeCopy_SmallString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]byte, len(str1)+len(bs))
		copy(s, str1)
		copy(s[len(str1):], bs)
		r2 = s
	}
}

var r1b []byte

func Benchmark_ConvertAppend_LargeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r1b = append([]byte(str2), bs...)
	}
}

var r2b []byte

func Benchmark_MakeCopy_LargeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]byte, len(str2)+len(bs))
		copy(s, str2)
		copy(s[len(str2):], bs)
		r2b = s
	}
}
