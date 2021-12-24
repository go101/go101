package bytes

import "testing"

type K struct {
	Dir, File string
}

var mk = make(map[K]struct{})
var ma = make(map[[2]string]struct{})
var ms = make(map[string]struct{})

var keyparts = []string{
	"docs", "aaa",
	"pictures", "bbb",
	"downloads", "ccc",
}

func fk(a, b string) {
	mk[K{a, b}] = struct{}{}
}

func fa(a, b string) {
	ma[[2]string{a, b}] = struct{}{}
}

func fs(a, b string) {
	ms[a+"/"+b] = struct{}{}
}

func Benchmark_struct_key(b1 *testing.B) {
	for i := 0; i < b1.N; i++ {
		for i := 0; i < len(keyparts); i += 2 {
			fk(keyparts[i], keyparts[i+1])
		}
		for key := range mk {
			delete(mk, key)
		}
	}
}

func Benchmark_array_key(b1 *testing.B) {
	for i := 0; i < b1.N; i++ {
		for i := 0; i < len(keyparts); i += 2 {
			fa(keyparts[i], keyparts[i+1])
		}
		for key := range ma {
			delete(ma, key)
		}
	}
}

func Benchmark_string_key(b1 *testing.B) {
	for i := 0; i < b1.N; i++ {
		for i := 0; i < len(keyparts); i += 2 {
			fs(keyparts[i], keyparts[i+1])
		}
		for key := range ms {
			delete(ms, key)
		}
	}
}
