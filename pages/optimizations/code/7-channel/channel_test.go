package channels

import "testing"

var vx int
var vy string

func Benchmark_TwoChannels(b *testing.B) {
	var x = make(chan int)
	var y = make(chan string)
	go func() {
		for {
			x <- 1
		}
	}()
	go func() {
		for {
			y <- "hello"
		}
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		select {
		case vx = <-x:
		case vy = <-y:
		}
	}
}

func Benchmark_OneChannel_Interface(b *testing.B) {
	var x = make(chan interface{})
	go func() {
		for {
			x <- 1
		}
	}()
	go func() {
		for {
			x <- "hello"
		}
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		select {
		case v := <-x:
			switch v := v.(type) {
			case int:
				vx = v
			case string:
				vy = v
			}
		}
	}
}

func Benchmark_OneChannel_Struct(b *testing.B) {
	type T struct {
		x int
		y string
	}
	var x = make(chan T)
	go func() {
		for {
			x <- T{x: 1}
		}
	}()
	go func() {
		for {
			x <- T{y: "hello"}
		}
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := <-x
		if v.y != "" {
			vy = v.y
		} else {
			vx = v.x
		}
	}
}
