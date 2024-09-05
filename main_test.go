package main

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func BenchmarkStr(b *testing.B) {
	file, err := os.Open("be.txt")
	all, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		var offset int = 0
		for offset < len(all) {

			bu := all[offset : offset+10]

			_ = bu
			offset += 10
		}

	}
}
func Benchmark_with_Buffer(b *testing.B) {
	file, err := os.Open("be.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReaderSize(file, 10000)
	for i := 0; i < b.N; i++ {
		for {

			buf := make([]byte, 10)
			_, err = reader.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					panic(err)
				}
			}

		}

	}
}
