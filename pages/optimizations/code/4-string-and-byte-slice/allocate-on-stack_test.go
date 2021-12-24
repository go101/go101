package bytes

import "testing"
import "bytes"
import "strings"

var s = "111111111111111111111111111+11111111111111111111+11111111111111111"
var s2 = s + s
var r string

func ReplaceByte1(str string, from, to byte) string {
	return strings.ReplaceAll(str, string(from), string(to))
}

func ReplaceByte2(str string, from, to byte) string {
	bs := []byte(str)
	return string(bytes.ReplaceAll(bs, []byte{from}, []byte{to}))
}

func ReplaceByte3(str string, from, to byte) string {
	bs := []byte(str)
	for s := bs; len(s) > 0; {
		if i := bytes.IndexByte(s, from); i >= 0 {
			s[i] = to
			s = s[i+1:]
		} else {
			break
		}
	}
	return string(bs)
}

func ReplaceByte4(str string, from, to byte) string {
	var bs []byte
	if len(str) > 128 {
		bs = []byte(str)
	} else {
		var a [128]byte
		bs = a[:len(str)]
		copy(bs, str)
	}
	for s := bs; len(s) > 0; {
		if i := bytes.IndexByte(s, from); i >= 0 {
			s[i] = to
			s = s[i+1:]
		} else {
			break
		}
	}
	return string(bs)
}

func Benchmark_ReplaceByte_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = ReplaceByte1(s, '+', ' ')
		r = ReplaceByte1(s2, '+', ' ')
	}
}

func Benchmark_ReplaceByte_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = ReplaceByte2(s, '+', ' ')
		r = ReplaceByte2(s2, '+', ' ')
	}
}

func Benchmark_ReplaceByte_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = ReplaceByte3(s, '+', ' ')
		r = ReplaceByte3(s2, '+', ' ')
	}
}

func Benchmark_ReplaceByte_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = ReplaceByte4(s, '+', ' ')
		r = ReplaceByte4(s2, '+', ' ')
	}
}
