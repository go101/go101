package functions

import "testing"

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

func bar(n int) {
	switch {
	case indexTable[n % 10]:
		// do something 1
	default:
		// do something 2
	}
}

var values []int

func init() {
	values = make([]int, 1000)
	var k = 0
	for i := 0; i < 10; i++ {
		n := (i + 1) * 10
		for n > 0 {
			values[k] = i
			k++
			n--
		}
	}
	values = values[:k]
}

func Benchmark_foo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range values {
			foo(v)
		}
	}
}

func Benchmark_bar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range values {
			bar(v)
		}
	}
}

