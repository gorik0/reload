package main

import (
	"os"
	"testing"
)

func Benchmark_Scanenr(b *testing.B) {
	file, err := os.Open("gorik.txt")
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		_, err := makeWordsFromFile_Scanenr(file)
		if err != nil {
			panic(err)

		}

	}
}
func Benchmark_Rural(b *testing.B) {
	file, err := os.Open("gorik.txt")
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		_, err := makeWordsFromFile_Rural(file)
		if err != nil {
			panic(err)

		}

	}
}
func Benchmark_String(b *testing.B) {
	file, err := os.Open("gorik.txt")
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		_, err := makeWordsFromFile_String(file)
		if err != nil {
			panic(err)

		}

	}
}
func Benchmark_Buffered(b *testing.B) {
	file, err := os.Open("gorik.txt")
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		_, err := makeWordsFromFile_Buffered(file)
		if err != nil {
			panic(err)

		}

	}
}
