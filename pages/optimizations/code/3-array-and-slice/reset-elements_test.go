package main

import (
	"testing"
	"fmt"
)

type A [1024]byte

var a A

func Benchmark_Reset_ArrayPtr(b *testing.B) {
	var p = &a
	
	for i := 0; i < b.N; i++ {
		for i := range p {
			p[i] = 0
		}
	}
}

func Benchmark_Reset_ArrayPtr_b(b *testing.B) {
	var p = &a
	
	for i := 0; i < b.N; i++ {
		*p = A{}
	}
}

func Benchmark_Reset_Array_memclr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := range a {
			a[i] = 0
		}
	}
}

func Benchmark_Reset_Array_assignment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a = A{}
	}
}

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
