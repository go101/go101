package bce

import (
	"math/rand"
	"testing"
	"time"
)

var N = 1 << 13
var as = make([][4]uint32, N)
var bs = make([]byte, 4*4*N)

func init() {
	rand.Seed(time.Now().UnixNano())
	s := as
	for i := range s {
		s[i] = [4]uint32{
			uint32(rand.Int63()),
			uint32(rand.Int63()),
			uint32(rand.Int63()),
			uint32(rand.Int63()),
		}
	}
}

func g(b []byte, v uint32) {
	b[3] = byte(v) // Found IsInBounds (line 4)
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
}

func f8a(s []byte, d [4]uint32) {
	g(s[0:], d[0])  // Found IsInBounds
	g(s[4:], d[1])  // Found IsInBounds
	g(s[8:], d[2])  // Found IsInBounds+IsSliceInBounds
	g(s[12:], d[3]) // Found IsInBounds+IsSliceInBounds
}

func f8a1(s []byte, d [4]uint32) {
	g(s[0:4], d[0])
	g(s[4:8], d[1])
	g(s[8:12], d[2])
	g(s[12:16], d[3])
}

func f8a2(s []byte, d [4]uint32) {
	g(s[0:4:4], d[0])
	g(s[4:8:8], d[1])
	g(s[8:12:12], d[2])
	g(s[12:16:16], d[3])
}

func f8b(s []byte, d [4]uint32) {
	s = s[:16:16] // Found IsSliceInBounds

	g(s[0:], d[0])
	g(s[4:], d[1])
	g(s[8:], d[2])
	g(s[12:], d[3])
}

func f8b1(s []byte, d [4]uint32) {
	s = s[:16:16] // Found IsSliceInBounds

	g(s[0:4], d[0])
	g(s[4:8], d[1])
	g(s[8:12], d[2])
	g(s[12:16], d[3])
}

func f8b2(s []byte, d [4]uint32) {
	s = s[:16:16] // Found IsSliceInBounds

	g(s[0:4:4], d[0])
	g(s[4:8:8], d[1])
	g(s[8:12:12], d[2])
	g(s[12:16:16], d[3])
}

func Benchmark_f8a(b *testing.B) {
	s := as
	for i := 0; i < b.N; i++ {
		k := 0
		for j := range s {
			bs := bs[k:]
			f8a(bs, s[j])
			k += 16
		}
	}
}

func Benchmark_f8a1(b *testing.B) {
	s := as
	for i := 0; i < b.N; i++ {
		k := 0
		for j := range s {
			bs := bs[k:]
			f8a1(bs, s[j])
			k += 16
		}
	}
}

func Benchmark_f8a2(b *testing.B) {
	s := as
	for i := 0; i < b.N; i++ {
		k := 0
		for j := range s {
			bs := bs[k:]
			f8a2(bs, s[j])
			k += 16
		}
	}
}

func Benchmark_f8b(b *testing.B) {
	s := as
	for i := 0; i < b.N; i++ {
		k := 0
		for j := range s {
			bs := bs[k:]
			f8b(bs, s[j])
			k += 16
		}
	}
}

func Benchmark_f8b1(b *testing.B) {
	s := as
	for i := 0; i < b.N; i++ {
		k := 0
		for j := range s {
			bs := bs[k:]
			f8b1(bs, s[j])
			k += 16
		}
	}
}

func Benchmark_f8b2(b *testing.B) {
	s := as
	for i := 0; i < b.N; i++ {
		k := 0
		for j := range s {
			bs := bs[k:]
			f8b2(bs, s[j])
			k += 16
		}
	}
}
