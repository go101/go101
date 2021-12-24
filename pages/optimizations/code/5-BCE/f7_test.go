package bce

import (
	"math/rand"
	"testing"
	"time"
)

var N = 1 << 13
var s = make([]byte, N)
var r = make([]byte, N/4)

func init() {
	rand.Seed(time.Now().UnixNano())
	s := s
	for i := range s {
		s[i] = byte(rand.Intn(256))
	}
}

//go:noinline
func f7a(rs []byte, bs []byte) {
	for i, j := 0, 0; i < len(bs)-3; i += 4 {
		rs[j] = bs[i+3] ^ bs[i+2] ^ bs[i+1] ^ bs[i]
		/*
			0x0032 00050 (f7_test.go:24)	LEAQ	3(CX), R8
			0x0036 00054 (f7_test.go:24)	CMPQ	SI, R8
			0x0039 00057 (f7_test.go:24)	JLS	159
			0x003b 00059 (f7_test.go:24)	MOVBLZX	3(CX)(DI*1), R8
			0x0041 00065 (f7_test.go:24)	LEAQ	2(CX), R9
			0x0045 00069 (f7_test.go:24)	CMPQ	SI, R9
			0x0048 00072 (f7_test.go:24)	JLS	148
			0x004a 00074 (f7_test.go:24)	MOVBLZX	2(CX)(DI*1), R9
			0x0050 00080 (f7_test.go:24)	XORL	R8, R9
			0x0053 00083 (f7_test.go:24)	LEAQ	1(CX), R8
			0x0057 00087 (f7_test.go:24)	CMPQ	SI, R8
			0x005a 00090 (f7_test.go:24)	JLS	137
			0x005c 00092 (f7_test.go:24)	MOVBLZX	1(CX)(DI*1), R8
			0x0062 00098 (f7_test.go:24)	XORL	R9, R8
			0x0065 00101 (f7_test.go:24)	MOVBLZX	(DI)(CX*1), R9
			0x006a 00106 (f7_test.go:24)	XORL	R8, R9
			0x006d 00109 (f7_test.go:24)	CMPQ	BX, DX
			0x0070 00112 (f7_test.go:24)	JHI	30

		*/
		j++
	}
}

//go:noinline
func f7b(rs []byte, bs []byte) {
	for i, j := 0, 0; i < len(bs)-3; i += 4 {
		bs := bs[i : i+4] // Found IsSliceInBounds
		rs[j] = bs[3] ^ bs[2] ^ bs[1] ^ bs[0]
		/*
			0x0031 00049 (f7_test.go:53)	LEAQ	4(CX), R9
			0x0035 00053 (f7_test.go:53)	CMPQ	R8, R9
			0x0038 00056 (f7_test.go:53)	JCS	147
			0x003a 00058 (f7_test.go:53)	CMPQ	CX, R9
			0x003d 00061 (f7_test.go:53)	JHI	136
			0x003f 00063 (f7_test.go:53)	MOVQ	CX, R10
			0x0042 00066 (f7_test.go:53)	SUBQ	R8, CX
			0x0045 00069 (f7_test.go:53)	SARQ	$63, CX
			0x0049 00073 (f7_test.go:53)	ANDQ	CX, R10
			0x004c 00076 (f7_test.go:54)	MOVBLZX	3(DI)(R10*1), R11
			0x0052 00082 (f7_test.go:54)	MOVBLZX	2(DI)(R10*1), R12
			0x0058 00088 (f7_test.go:54)	XORL	R11, R12
			0x005b 00091 (f7_test.go:54)	MOVBLZX	1(DI)(R10*1), R11
			0x0061 00097 (f7_test.go:54)	XORL	R12, R11
			0x0064 00100 (f7_test.go:54)	MOVBLZX	(DI)(R10*1), R10
			0x0069 00105 (f7_test.go:54)	XORL	R11, R10
			0x006c 00108 (f7_test.go:54)	CMPQ	BX, DX
			0x006f 00111 (f7_test.go:54)	JHI	30

		*/
		j++
	}
}

//go:noinline
func f7c(rs []byte, bs []byte) {
	for i, j := 0, 0; i < len(bs)-3; i += 4 {
		bs := bs[i : i+4 : i+4] // Found IsSliceInBounds
		rs[j] = bs[3] ^ bs[2] ^ bs[1] ^ bs[0]
		/*
			0x0031 00049 (f7_test.go:83)	LEAQ	4(CX), R9
			0x0035 00053 (f7_test.go:83)	CMPQ	R8, R9
			0x0038 00056 (f7_test.go:83)	JCS	135
			0x003a 00058 (f7_test.go:83)	CMPQ	CX, R9
			0x003d 00061 (f7_test.go:83)	JHI	124
			0x003f 00063 (f7_test.go:84)	MOVBLZX	3(DI)(CX*1), R10
			0x0045 00069 (f7_test.go:84)	MOVBLZX	2(DI)(CX*1), R11
			0x004b 00075 (f7_test.go:84)	XORL	R10, R11
			0x004e 00078 (f7_test.go:84)	MOVBLZX	1(DI)(CX*1), R10
			0x0054 00084 (f7_test.go:84)	XORL	R11, R10
			0x0057 00087 (f7_test.go:84)	MOVBLZX	(DI)(CX*1), R11
			0x005c 00092 (f7_test.go:84)	XORL	R10, R11
			0x005f 00095 (f7_test.go:84)	NOP
			0x0060 00096 (f7_test.go:84)	CMPQ	BX, DX
			0x0063 00099 (f7_test.go:84)	JHI	30

		*/
		j++
	}
}

func Benchmark_f7a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f7a(r, s)
	}
}

func Benchmark_f7b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f7b(r, s)
	}
}

func Benchmark_f7c(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f7c(r, s)
	}
}
