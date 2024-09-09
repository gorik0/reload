package main

import "testing"

var nums = GenNumbers(1000000)

func Benchmark___add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(nums)
	}
}
func BenchmarkName___goroutineADD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add_HARD(5, nums)
	}
}
