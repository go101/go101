package main

import (
	"testing"
	"fmt"
)

func Benchmark_ResetElements(b *testing.B) {
	const MaxN = 1024 * 1024 * 32
	var array [MaxN]byte
	
	for n := MaxN >> 1; n >= 2; n >>= 1 {
		b.Run(fmt.Sprintf("reset_%d", n), func(b *testing.B){
			for i := 0; i < b.N; i++ {
				s := array[:n+1]
				for k := 0; k < len(s); k++ {
					s[k] = 0
				}
			}
		})
		b.Run(fmt.Sprintf("memclr_%d", n), func(b *testing.B){
			for i := 0; i < b.N; i++ {
				s := array[:n+1]
				for k, _ := range s {
					s[k] = 0
				}
			}
		})
	}
}
