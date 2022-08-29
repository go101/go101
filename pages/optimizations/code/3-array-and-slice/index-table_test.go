package arrays

// https://github.com/golang/go/issues/27857

import "testing"

//go:noinline
func foo(n int) {
	switch n % 10 {
	case 1, 2, 6, 7, 9:
		// do something 1
	default:
		// do something 2
	}
}

var indexTable = [10]bool {
	1: true, 2: true, 6: true, 7: true, 9: true,
}

//go:noinline
func bar(n int) {
	switch {
	case indexTable[n % 10]:
		// do something 1
	default:
		// do something 2
	}
}

func Benchmark_foo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := 0; k < 100; k++ {
			foo(i+k)
		}
	}
}

func Benchmark_bar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := 0; k < 100; k++ {
			bar(i+k)
		}
	}
}

func Benchmark_bar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := 0; k < 100; k++ {
			bar(i+k)
		}
	}
}

func Benchmark_foo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := 0; k < 100; k++ {
			foo(i+k)
		}
	}
}

