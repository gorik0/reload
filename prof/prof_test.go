package main

import "testing"

//var hardString = makeString()

func Benchmark____mapString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapString()
	}
}
func Benchmark____mapInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapInt()
	}
}
