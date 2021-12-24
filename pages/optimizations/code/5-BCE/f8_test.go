package bce

import (
	"testing"
)

var N = 1 << 13
var ss = make([][]byte, 256)

func init() {
	for i := range ss {
		ss[i] = make([]byte, N)
	}
}

//go:noinline
func f8x(s []byte) {
	var n byte
	for i := 0; i < len(s)-3; i += 4 {
		s[i+3] = n // Found IsInBounds
		s[i+2] = n // Found IsInBounds
		s[i+1] = n // Found IsInBounds
		s[i] = n
		n++
	}
}

//go:noinline
func f8y1(s []byte) {
	var n byte
	for i := 0; i < len(s)-3; i += 4 {
		s2 := s[i:]
		s2[3] = n // Found IsInBounds
		s2[2] = n
		s2[1] = n
		s2[0] = n
		n++
	}
}

func f8y2(s []byte) {
	var n byte
	for i := 0; i < len(s)-3; i += 4 {
		s2 := s[i : i+4]
		s2[3] = n // Found IsInBounds
		s2[2] = n
		s2[1] = n
		s2[0] = n
		n++
	}
}

func f8y3(s []byte) {
	var n byte
	for i := 0; i < len(s)-3; i += 4 {
		s2 := s[i : i+4 : i+4]
		s2[3] = n // Found IsInBounds
		s2[2] = n
		s2[1] = n
		s2[0] = n
		n++
	}
}

//go:noinline
func f8z(s []byte) {
	var n byte
	for i := 0; len(s) >= 4; i += 4 {
		s[3] = n
		s[2] = n
		s[1] = n
		s[0] = n
		s = s[4:]
		n++
	}
}

func Benchmark_f8x(b *testing.B) {
	ss := ss
	for i := 0; i < b.N; i++ {
		for _, s := range ss {
			f8x(s)
		}
	}
}

func Benchmark_f8y1(b *testing.B) {
	ss := ss
	for i := 0; i < b.N; i++ {
		for _, s := range ss {
			f8y1(s)
		}
	}
}

func Benchmark_f8y2(b *testing.B) {
	ss := ss
	for i := 0; i < b.N; i++ {
		for _, s := range ss {
			f8y2(s)
		}
	}
}

func Benchmark_f8y3(b *testing.B) {
	ss := ss
	for i := 0; i < b.N; i++ {
		for _, s := range ss {
			f8y3(s)
		}
	}
}

func Benchmark_f8z(b *testing.B) {
	ss := ss
	for i := 0; i < b.N; i++ {
		for _, s := range ss {
			f8z(s)
		}
	}
}
