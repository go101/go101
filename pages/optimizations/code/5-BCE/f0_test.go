package bce

import (
	"math/rand"
	"testing"
	"time"
)

const N = 1 << 13

var s = make([]byte, N*16)
var r = make([]byte, N*4)

func init() {
	rand.Seed(time.Now().UnixNano())
	s := s
	for i := range s {
		s[i] = byte(rand.Intn(256))
	}

	var x = make([]byte, N*4)
	var y = make([]byte, N*4)

	for k := 0; k < N; k++ {
		*(*[4]byte)(x[k*4:]) = f0a(*(*[16]byte)(s[k*16:]))
	}
	for k := 0; k < N; k++ {
		*(*[4]byte)(y[k*4:]) = f0b(*(*[16]byte)(s[k*16:]))
	}
	var allAreZeros = false
	for i, v := range x {
		if v != y[i] {
			panic("x[i] != y[i]")
		}
		if v != 0 {
			allAreZeros = false
		}
	}
	if allAreZeros {
		panic("all are zeros")
	}
}

func f0a(x [16]byte) (r [4]byte) {
	for i := 0; i < 4; i++ {
		r[i] =
			x[i*4+3] ^ // Found IsInBounds
				x[i*4+2] ^ // Found IsInBounds
				x[i*4+1] ^ // Found IsInBounds
				x[i*4] // Found IsInBounds
	}
	return
}

func f0b(x [16]byte) (r [4]byte) {
	r[0] = x[3] ^ x[2] ^ x[1] ^ x[0]
	r[1] = x[7] ^ x[6] ^ x[5] ^ x[4]
	r[2] = x[11] ^ x[10] ^ x[9] ^ x[8]
	r[3] = x[15] ^ x[14] ^ x[13] ^ x[12]
	return
}

func Benchmark_f0a(b *testing.B) {
	s, r := s, r
	for i := 0; i < b.N; i++ {
		for k := 0; k < N; k++ {
			*(*[4]byte)(r[k*4:]) = f0a(*(*[16]byte)(s[k*16:]))
		}
	}
}

func Benchmark_f0b(b *testing.B) {
	s, r := s, r
	for i := 0; i < b.N; i++ {
		for k := 0; k < N; k++ {
			*(*[4]byte)(r[k*4:]) = f0b(*(*[16]byte)(s[k*16:]))
		}
	}
}
