package allocations

import "testing"

const N = 821

type Book struct {
	Title  string
	Author string
	Pages  int
}

func CreateBooksOnOneLargeBlock(n int) []*Book {
	type _ int // avoid being inline

	books := make([]Book, n)
	pbooks := make([]*Book, n)
	for i := range pbooks {
		pbooks[i] = &books[i]
	}
	return pbooks
}

func CreateBooksOnManySmallBlocks(n int) []*Book {
	type _ int // avoid being inline

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
