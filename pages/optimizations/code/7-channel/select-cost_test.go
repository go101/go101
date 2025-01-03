package channels

import "testing"

var ch1 = make(chan struct{}, 1)
var ch2 = make(chan struct{}, 1)

func Benchmark_Select_OneCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case ch1 <- struct{}{}:
			<-ch1
		}
	}
}

func Benchmark_Select_TwoCases(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case ch1 <- struct{}{}:
			<-ch1
		case ch2 <- struct{}{}:
			<-ch2
		}
	}
}

var ch3 = make(chan struct{}, 1)
var ch4 = make(chan struct{}, 1)
var ch5 = make(chan struct{}, 1)

func Benchmark_Select_TwoCases_Plus_TryReceive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case <-chan int(nil):
		default:
		}

		select {
		case ch1 <- struct{}{}:
			<-ch1
		case ch2 <- struct{}{}:
			<-ch2
		}
	}
}

func Benchmark_Select_TwoCases_Plus_TrySent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case chan int(nil) <- 1:
		default:
		}

		select {
		case ch1 <- struct{}{}:
			<-ch1
		case ch2 <- struct{}{}:
			<-ch2
		}
	}
}

func Benchmark_Select_ThreeCases(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case ch1 <- struct{}{}:
			<-ch1
		case ch2 <- struct{}{}:
			<-ch2
		case ch3 <- struct{}{}:
			<-ch3
		}
	}
}

func Benchmark_Select_FourCases(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case ch1 <- struct{}{}:
			<-ch1
		case ch2 <- struct{}{}:
			<-ch2
		case ch3 <- struct{}{}:
			<-ch3
		case ch4 <- struct{}{}:
			<-ch4
		}
	}
}

func Benchmark_Select_FiveCases(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case ch1 <- struct{}{}:
			<-ch1
		case ch2 <- struct{}{}:
			<-ch2
		case ch3 <- struct{}{}:
			<-ch3
		case ch4 <- struct{}{}:
			<-ch4
		case ch5 <- struct{}{}:
			<-ch5
		}
	}
}
