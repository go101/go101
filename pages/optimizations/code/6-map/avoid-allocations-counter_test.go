package maps

import "testing"

var wordCounterA = make(map[string]int)
var wordCounterB = make(map[string]*int)
var key = make([]byte, 64)

func IncA(w []byte) {
	wordCounterA[string(w)]++
}

func IncB(w []byte) {
	p := wordCounterB[string(w)]
	if p == nil {
		p = new(int)
		wordCounterB[string(w)] = p
	}
	*p++
}

func Benchmark_A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := range key {
			IncA(key[:i])
		}
	}
}

func Benchmark_B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := range key {
			IncB(key[:i])
		}
	}
}

var wordIndexes = make(map[string]int)
var wordCounters []int

func IncC(w []byte) {
	if i, ok := wordIndexes[string(w)]; ok {
		wordCounters[i]++
	} else {
		wordIndexes[string(w)] = len(wordCounters)
		wordCounters = append(wordCounters, 1)
	}
}

func Benchmark_C(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := range key {
			IncC(key[:i])
		}
	}
}
