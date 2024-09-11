package main

import (
	"testing"
)

func Benchmark____SplitByLines___Buffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitByLines___Buffer__BUFIO("../" + fileName)
	}
}
func Benchmark____SplitByLines___Split(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitByLines___Split("../" + fileName)
	}
}
func Benchmark____SplitByLines___StringScanner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitByLines___StringScanner("../" + fileName)
	}
}
func BenchmarkSplitByLines___Buffer__BYTES(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitByLines___Buffer__BYTES("../" + fileName)
	}
}
