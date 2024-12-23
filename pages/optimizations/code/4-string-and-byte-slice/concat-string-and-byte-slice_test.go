package bytes

import "bytes"
import "testing"

var K = 1000
var N = 1 * K
var M = 2 * K

var str = string(make([]byte, N))
var bs = make([]byte, M)

var x []byte

func Benchmark_makecopy(b *testing.B) {
	for b.Loop() {
		var newByteSlice = make([]byte, len(str)+len(bs))
		copy(newByteSlice, str)
		copy(newByteSlice[len(str):], bs)
		x = newByteSlice
	}
}

func Benchmark_append(b *testing.B) {
	for b.Loop() {
		x = append([]byte(str), bs...)
	}
}

func Benchmark_Join(b *testing.B) {
	for b.Loop() {
		x = bytes.Join([][]byte{[]byte(str), bs}, nil)
	}
}

func Benchmark_Join2(b *testing.B) {
	for b.Loop() {
		x = bytes.Join([][]byte{[]byte(str), nil}, bs)
	}
}

func Benchmark_Convert(b *testing.B) {
	for b.Loop() {
		x = []byte(" " + string(str) + string(bs))[1:]
	}
}

