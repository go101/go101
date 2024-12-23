package bytes

import "testing"
import "bytes"

var x = bytes.Repeat([]byte("x"), 33)
var y = bytes.Repeat([]byte("y"), 33)
var z = bytes.Repeat([]byte("z"), 2)

func verbose(x, y, z []byte){
	switch {
	case string(x) == string(y):
		// do something
	case string(x) == string(z):
		// do something
	}
}

func clean(x, y, z []byte){
	switch string(x) {
	case string(y):
		// do something
	case string(z):
		// do something
	}
}

func Benchmark_verbose(b *testing.B) {
	for range b.N {
		verbose(x, y, z)
	}
}

func Benchmark_clean(b *testing.B) {
	for range b.N {
		clean(x, y, z)
	}
}



