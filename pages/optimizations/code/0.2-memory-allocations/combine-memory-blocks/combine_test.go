package allocations

import "testing"

const N = 100 // 821

type Book struct {
	Title  string
	Author string
	Pages  int
}

//go:noinline
func CreateBooksOnOneLargeBlock(n int) []*Book {
	books := make([]Book, n)
	pbooks := make([]*Book, n)
	for i := range pbooks {
		pbooks[i] = &books[i]
	}
	return pbooks
}

//go:noinline
func CreateBooksOnManySmallBlocks(n int) []*Book {
	books := make([]*Book, n)
	for i := range books {
		books[i] = new(Book)
	}
	return books
}

func Benchmark_CreateOnOneLargeBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CreateBooksOnOneLargeBlock(N)
	}
}

func Benchmark_CreateOnManySmallBlocks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CreateBooksOnManySmallBlocks(N)
	}
}
