package main

import "sync"
import "testing"

func Benchmark_Slice(b *testing.B) {
	var pool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 256)
		},
	}
	var bs []byte
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		bs = pool.Get().([]byte)
		bs[0] = 123
		pool.Put(bs)
	}
}

func Benchmark_SlicePointer(b *testing.B) {
	var pool = sync.Pool{
		New: func() interface{} {
			//return make([]byte, 256)
			return &[]byte{255, 0}
		},
	}
	var bs []byte
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		bs = *pool.Get().(*[]byte)
		bs[0] = 123
		pool.Put(&bs)
	}
}

func Benchmark_CustomPool_NoSync(b *testing.B) {
	type Node struct {
		s []byte
		n *Node
	}
	var head *Node
	var Get = func() *Node {
		if head == nil {
			return &Node{
				s: make([]byte, 256),
			}
		}
		n := head
		head = head.n
		return n
	}
	var Put = func(n *Node) {
		n.n = head
		head = n
	}
	var bs []byte
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		node := Get()
		bs = node.s
		bs[0] = 123
		Put(node)
	}
}

func Benchmark_CustomPool_Synced(b *testing.B) {
	type Node struct {
		s []byte
		n *Node
	}
	var m sync.Mutex
	var head *Node
	var Get = func() *Node {
		m.Lock()
		defer m.Unlock()
		if head == nil {
			return &Node{
				s: make([]byte, 256),
			}
		}
		n := head
		head = head.n
		return n
	}
	var Put = func(n *Node) {
		m.Lock()
		defer m.Unlock()
		n.n = head
		head = n
	}
	var bs []byte
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		node := Get()
		bs = node.s
		bs[0] = 123
		Put(node)
	}
}

/*
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: a.y/bench/pool
cpu: Intel(R) Core(TM) i5-4210U CPU @ 1.70GHz
Benchmark_Slice-4               	15991357	        73.61 ns/op	      24 B/op	       1 allocs/op
Benchmark_SlicePointer-4        	44608159	        25.91 ns/op	       0 B/op	       0 allocs/op
Benchmark_CustomPool_NoSync-4   	289593338	         4.135 ns/op	       0 B/op	       0 allocs/op
Benchmark_CustomPool_Synced-4   	28623903	        41.78 ns/op	       0 B/op	       0 allocs/op
*/
