package main

import "testing"

func BenchmarkMi___SINGLE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mi(1)
	}
}
func BenchmarkMi___MULTI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mi(12)
	}
}
func BenchmarkMi___MANY(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mi(100)
	}
}
