package channels

import "testing"

var c = make(chan struct{})

func Benchmark_TryReceive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case <-c:
		default:
		}
	}
}

func Benchmark_TrySend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case c <- struct{}{}:
		default:
		}
	}
}
