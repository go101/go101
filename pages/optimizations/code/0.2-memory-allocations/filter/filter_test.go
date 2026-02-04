package allocations

import "testing"

func buildOriginalData() []int {
	s := make([]int, 1024)
	for i := range s {
		s[i] = i
	}
	return s
}

func check(v int) bool {
	return v%2 == 0
}

func FilterOneAllocation(data []int) []int {
	var r = make([]int, 0, len(data))
	for _, v := range data {
		if check(v) {
			r = append(r, v)
		}
	}
	return r
}

func FilterNoAllocations(data []int) []int {
	var k = 0
	for i, v := range data {
		if check(v) {
			data[i] = data[k]
			data[k] = v
			k++
		}
	}
	return data[:k]
}

func Benchmark_FilterOneAllocation(b *testing.B) {
	data := buildOriginalData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FilterOneAllocation(data)
	}
}

func Benchmark_FilterNoAllocations(b *testing.B) {
	data := buildOriginalData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FilterNoAllocations(data)
	}
}

/*

import "reflect"

func init() {
	s1 := FilterNoAllocations(buildOriginalData())
	s2 := FilterOneAllocation(buildOriginalData())
	s3 := FilterOneSmallerAllocation(buildOriginalData())
	if !reflect.DeepEqual(s1, s2) {
		panic("s1 != s2")
	}
	if !reflect.DeepEqual(s1, s3) {
		panic("s1 != s3")
	}
}


func FilterOneSmallerAllocation(data []int) []int {
	var n = 0
	for _, v := range data {
		if check(v) {
			n++
		}
	}
	var r = make([]int, 0, n)
	for _, v := range data {
		if check(v) {
			r = append(r, v)
		}
	}
	return r
}

func Benchmark_FilterOneSmallerAllocation(b *testing.B) {
	data := buildOriginalData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FilterOneSmallerAllocation(data)
	}
}

/**/
