package maps

import "testing"

type byteArray interface {
	byte | int16 | [1]byte | [2]byte | [3]byte | [4]byte | [5]byte | [6]byte | [7]byte | [8]byte | [15]byte  | [16]byte  | [17]byte | [31]byte | [32]byte | [33]byte
}

func benchmarkMaps[T byteArray](b *testing.B) {
	m := map[T]int{}
	var t T
	for i := 0; i < b.N; i++ {
		m[t]++
	}
}

func BenchmarkMaps(b *testing.B) {
	b.Run("byte", benchmarkMaps[byte])
	b.Run("int16", benchmarkMaps[int16])
	b.Run("1", benchmarkMaps[[1]byte])
	b.Run("2", benchmarkMaps[[2]byte])
	b.Run("3", benchmarkMaps[[3]byte])
	b.Run("4", benchmarkMaps[[4]byte])
	b.Run("5", benchmarkMaps[[5]byte])
	b.Run("6", benchmarkMaps[[6]byte])
	b.Run("7", benchmarkMaps[[7]byte])
	b.Run("8", benchmarkMaps[[8]byte])
	b.Run("15", benchmarkMaps[[15]byte])
	b.Run("16", benchmarkMaps[[16]byte])
	b.Run("17", benchmarkMaps[[17]byte])
	b.Run("31", benchmarkMaps[[31]byte])
	b.Run("32", benchmarkMaps[[32]byte])
	b.Run("33", benchmarkMaps[[33]byte])
}


