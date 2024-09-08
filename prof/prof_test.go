package prof

import "testing"

//var hardString = makeString()

func Benchmark___Asignment(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Assignment()
	}
}
func Benchmark___SimpleConcat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		SimpleConcat()
	}
}
func Benchmark___ThroughBuilder(b *testing.B) {

	for i := 0; i < b.N; i++ {
		ThroughBuilder()
	}
}
func Benchmark___ThroughBuffer(b *testing.B) {

	for i := 0; i < b.N; i++ {
		ThroughBuffer()
	}
}
