package bytes

import "testing"
import "strings"

func CompareStrings(x, y string) int {
	var one = 1
	if len(x) > len(y) {
		x, y = y, x
		one = -1
	}
	if len(x) <= len(y) { // BCE hint
		for i := 0; i < len(x); i++ {
			if x[i] != y[i] {
				if x[i] > y[i] {
					return one
				}
				return -one
			}
		}
	}
	if len(x) < len(y) {
		return -one
	}
	return 0
}

func init() {
	if CompareStrings(ss[0], ss[1]) != 1 {
		panic("CompareStrings(ss[0], ss[1]) != 1")
	}
	if CompareStrings(ss[1], ss[0]) != -1 {
		panic("CompareStrings(ss[1], ss[0]) != -1")
	}
	if CompareStrings(ss[0], ss[0]) != 0 {
		panic("CompareStrings(ss[0], ss[0]) != 0")
	}
	if CompareStrings(ss[1], ss[1]) != 0 {
		panic("CompareStrings(ss[1], ss[1]) != 0")
	}
}

var ss = []string{
	"Hello world! Go! Go! Go!",
	"Hello world! Go! Go! Go",
}
var r bool

func Benchmark_StringsCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range ss {
			for _, y := range ss {
				r = strings.Compare(x, y) == 0
			}
		}
	}
}

func Benchmark_Operator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range ss {
			for _, y := range ss {
				r = x == y
			}
		}
	}
}

func Benchmark_CompareStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range ss {
			for _, y := range ss {
				r = CompareStrings(x, y) == 0
			}
		}
	}
}
