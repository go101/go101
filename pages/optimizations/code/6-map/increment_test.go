package maps

import "testing"

var m = map[int]int{}

func Benchmark_increment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m[99]++
	}
}

func Benchmark_plusone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m[99] += 1
	}
}

func Benchmark_addition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m[99] = m[99] + 1
	}
}

var ms = map[string]string{}

func Benchmark_put_delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ms["hello world!"] = "Go"
		delete(ms, "hello world!")
	}
}

func Benchmark_self_plus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ms["hello world!"] = "Go"
		ms["hello world!"] += " language"
		delete(ms, "hello world!")
	}
}

func Benchmark_concatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ms["hello world!"] = "Go"
		ms["hello world!"] =
			ms["hello world!"] + " language"
		delete(ms, "hello world!")
	}
}
