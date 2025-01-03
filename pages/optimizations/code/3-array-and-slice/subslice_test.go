package arrays

// https://github.com/golang/go/issues/27857

import "testing"
import "math/rand"

// Making benchmarks in Go is harder and harder.
// Since Go 1.21, if N == 1 << 6, unxpected results will be got,
// Still not get the reason.

const N = 1 << 10
var s = make([]byte, N)
var r = make([]byte, N/4)

func init() {
	s := s
	for i := range s {
		s[i] = byte(rand.Intn(256))
	}
}

//go:noinline
func h(rs []byte, bs []byte) {
	for i, j := 0, 0; i < len(bs) - 3; i += 4 {
		s2 := bs[i:]
		rs[j] = s2[3] ^ s2[0]
		j++
	}
}

//go:noinline
func f(rs []byte, bs []byte) {
	for i, j := 0, 0; i < len(bs) - 3; i += 4 {
		s2 := bs[i:i+4]
		rs[j] = s2[3] ^ s2[0]
		j++
	}
}

//go:noinline
func g(rs []byte, bs []byte) {
	for i, j := 0, 0; i < len(bs) - 3; i += 4 {
		s2 := bs[i:i+4:i+4]
		rs[j] = s2[3] ^ s2[0]
		j++
	}
}

//go:noinline
func q(rs []byte, bs []byte) {
	for j := 0; len(bs) >= 4; j++ {
		rs[j] = bs[3] ^ bs[0]
		bs = bs[4:]
	}
}

func Benchmark_h(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h(r, s)
	}
}

func Benchmark_f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(r, s)
	}
}

func Benchmark_g(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g(r, s)
	}
}

func Benchmark_q(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q(r, s)
	}
}
