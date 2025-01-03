package bytes

import "testing"
import "bytes"
import "strings"
import "fmt"
import "unsafe"

const M, N, K = 12, 16, 32

var s1 = strings.Repeat("a", M)
var s2 = strings.Repeat("a", N)
var s3 = strings.Repeat("a", K)
var r string

func init() {
	println("======", M, N, K)
}

func Concat_WithPlus(a, b, c string) string {
	return a + b + c
}

func Concat_WithBuilder(ss ...string) string {
	var b strings.Builder
	var n = 0
	for _, s := range ss {
		n += len(s)
	}
	b.Grow(n)
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}

func Concat_WithBytes(ss ...string) string {
	if len(ss) == 0 {
		return ""
	}
	var n = 0
	for _, s := range ss {
		n += len(s)
	}
	var s = ss[0]
	var bs = make([]byte, n)
	copy(bs, s)
	bs = bs[:len(s)]
	for _, s := range ss[1:] {
		bs = append(bs, s...)
	}
	return string(bs)
}

func Concat_WithUnsafe(ss ...string) string {
	if len(ss) == 0 {
		return ""
	}
	var n = 0
	for _, s := range ss {
		n += len(s)
	}
	var s = ss[0]
	var bs = make([]byte, n)
	copy(bs, s)
	bs = bs[:len(s)]
	for _, s := range ss[1:] {
		bs = append(bs, s...)
	}
	return *(*string)(unsafe.Pointer(&bs))
}

func Concat_WithPrint(a, b, c string) string {
	return fmt.Sprint(a, b, c)
}

func Concat_WithBuffer(ss ...string) string {
	var b bytes.Buffer
	var n = 0
	for _, s := range ss {
		n += len(s)
	}
	b.Grow(n)
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}

var gb bytes.Buffer

func Concat_WithReusedBuffer(ss ...string) string {
	gb.Reset()
	var n = 0
	for _, s := range ss {
		n += len(s)
	}
	gb.Grow(n)
	for _, s := range ss {
		gb.WriteString(s)
	}
	return gb.String()
}

func Benchmark_Concat_WithPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Concat_WithPlus(s1, s2, s3)
	}
}

func Benchmark_Concat_WithBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Concat_WithBuilder(s1, s2, s3)
	}
}

func Benchmark_Concat_WithBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Concat_WithBytes(s1, s2, s3)
	}
}

func Benchmark_Concat_WithUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Concat_WithUnsafe(s1, s2, s3)
	}
}

func Benchmark_Concat_WithPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Concat_WithPrint(s1, s2, s3)
	}
}

func Benchmark_Concat_WithBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Concat_WithBuffer(s1, s2, s3)
	}
}

func Benchmark_Concat_WithReusedBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Concat_WithReusedBuffer(s1, s2, s3)
	}
}

func Benchmark_Concat_WithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = strings.Join([]string{s1, s2, s3}, "")
	}
}
