package channels

import (
	"sync"
	"sync/atomic"
	"testing"
)

var g int32

func Benchmark_NoSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g++
	}
}

func Benchmark_Atomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomic.AddInt32(&g, 1)
	}
}

var m sync.Mutex

func Benchmark_Mutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Lock()
		g++
		m.Unlock()
	}
}

var ch = make(chan struct{}, 1)

func Benchmark_Channel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
		g++
		<-ch
	}
}
