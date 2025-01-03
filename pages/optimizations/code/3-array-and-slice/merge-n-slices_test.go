package arrays

import "testing"
import "math/rand"
import "time"

var ss [][]byte
var r []byte

func init() {
	rand.Seed(time.Now().UnixNano())
	
	ss = make([][]byte, 6)
	for i := range ss {
		// Before Go 1.17, if the length of the merged result slice is not larger than 32768,
		// MergeN_MakeCopy is more performant, otherwise, MergeN_Orderless is more performant.
		const N = 10000
		if i == len(ss) - 1 {
			ss[i] = make([]byte, N * i * 5)
		} else {
			ss[i] = make([]byte, N * i)
		}
		for k := range ss[i] {
			ss[i][k] = byte(rand.Int())
		}
	}
	var n = 0
	for _, s := range ss {
		n += len(s)
	}
	s1 := MergeN_MakeCopy(ss...)
	s2 := MergeN_MakeAppend(ss...)
	s3 := MergeN_MakeCopyAppend(ss...)
	s4 := MergeN_Orderless(ss...)
	if len(s1) != n {
		panic("len(s1) != n")
	}
	if len(s2) != n {
		panic("len(s2) != n")
	}
	if len(s3) != n {
		panic("len(s3) != n")
	}
	if len(s4) != n {
		panic("len(s4) != n")
	}
	// If this if-block is enabled, then the
	// function MergeN_MakeCopyAppend will
	// consume less CPU in the benchmark results.
	//
	//if len(MergeN_Orderless(ss)) != n {
	//	panic("len(s4) != n")
	//}
	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			panic("s1[i] != s2[i]")
		}
		if s1[i] != s3[i] {
			panic("s1[i] != s3[i]")
		}
	}
}

func MergeN_MakeCopy(ss ...[]byte) []byte {
	var n = 0
	for _, s := range ss {
		n += len(s)
	}
	i, r := 0, make([]byte, n)
	for _, s := range ss {
		copy(r[i:], s)
		i += len(s)
	}
	return r
}

func MergeN_MakeAppend(ss ...[]byte) []byte {
	var n = 0
	for _, s := range ss {
		n += len(s)
	}
	r := make([]byte, 0, n)
	for _, s := range ss {
		r = append(r, s...)
	}
	return r
}

func MergeN_MakeCopyAppend(ss ...[]byte) []byte {
	var n = 0
	for _, s := range ss {
		n += len(s)
	}
	var s = ss[0]
	var r = make([]byte, n)
	copy(r, s)
	r = r[:len(s)]
	for _, s := range ss[1:] {
		r = append(r, s...)
	}
	return r
}

func MergeN_Orderless(ss ...[]byte) []byte {
	if len(ss) == 0 {
		return []byte{}
	}
	var k, max = 0, len(ss[0])
	var n = max
	for i := 1; i < len(ss); i++ {
		li := len(ss[i])
		n += li
		if li > max {
			k = i
			max = li
		}
	}
	
	if k > 0 {
		ss2 := make([][]byte, len(ss))
		copy(ss2, ss)
		ss2[k], ss2[0] = ss2[0], ss2[k]
		ss = ss2
	}
	var s = ss[0]
	var r = make([]byte, n)
	copy(r, s)
	r = r[:max]
	
	for _, s := range ss[1:] {
		r =append(r, s...)
	}
	return r
}

func Benchmark_MergeN_MakeCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = MergeN_MakeCopy(ss...)
	}
}

func Benchmark_MergeN_MakeAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = MergeN_MakeAppend(ss...)
	}
}

func Benchmark_MergeN_MakeCopyAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = MergeN_MakeCopyAppend(ss...)
	}
}

func Benchmark_MergeN_Orderless(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = MergeN_Orderless(ss...)
	}
}


