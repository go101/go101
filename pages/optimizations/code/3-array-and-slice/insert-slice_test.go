package arrays

import "testing"
import "math/rand"
import "time"

// Before Go 1.17, if the length of the final result slice is not larger than 32768,
// Insert1 is more performant than Insert2, otherwise, Insert2 is more performant.
var N = 16384 + 1 // 16384
var M = N
var k = N/2
var x1, x2 []byte
var r []byte

func init() {
	rand.Seed(time.Now().UnixNano())
	x1 = make([]byte, N)
	x2 = make([]byte, M)
	init := func(s []byte) {
		for i := range s {
			s[i] = byte(rand.Intn(256))
		}
	}
	init(x1)
	init(x2)
	
	a := Insert1(x1, k, x2)
	b := Insert1(x1, k, x2)
	c := Insert2(x1, k, x2)
	if len(a) != len(b) {
		panic("len(a) != len(b)")
	}
	if len(a) != len(c) {
		panic("len(a) != len(c)")
	}
	for i, v := range a {
		if v != b[i] {
			panic("a[i] != b[i]")
		}
		if v != c[i] {
			panic("a[i] != c[i]")
		}
	}
}

func Insert0(s []byte, k int, vs []byte) []byte {
	s2 := make([]byte, 0, len(s) + len(vs))
	s2 = append(s2, s[:k]...)
	s2 = append(s2, vs...)
	s2 = append(s2, s[k:]...)
	return s2
}

func Insert1(s []byte, k int, vs []byte) []byte {
	s2 := make([]byte, len(s) + len(vs))
	copy(s2, s[:k])
	copy(s2[k:], vs)
	copy(s2[k+len(vs):], s[k:])
	return s2
}

func Insert2(s []byte, k int, vs []byte) []byte {
	a := s[:k]
	s2 := make([]byte, len(s) + len(vs))
	copy(s2, a)
	copy(s2[len(a):], vs)
	copy(s2[len(a)+len(vs):], s[k:])
	return s2
}

func Insert3(s []byte, k int, vs []byte) []byte {
	return append(x1[:k:k], append(vs[:len(vs):len(vs)], x1[k:]...)...)
}

func Benchmark_Insert0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Insert0(x1, k, x2)
	}
}

func Benchmark_Insert1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Insert1(x1, k, x2)
	}
}

func Benchmark_Insert2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Insert2(x1, k, x2)
	}
}

func Benchmark_Insert3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = Insert3(x1, k, x2)
	}
}

