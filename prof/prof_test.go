package main

import (
	"testing"
)

//var hardString = makeString()

func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10000; i++ {
			makeRIOandCOUNT(1e4)

		}

	}
}
func BenchmarkBuilder___Pool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10000; i++ {
			makeRIOandCOUNT_____sync(1e4)

		}
	}
}
