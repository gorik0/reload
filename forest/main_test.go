package forest

import "testing"

func Benchmark___jso(b *testing.B) {
	for i := 0; i < b.N; i++ {

		jso()
	}
}
func Benchmark___jso___EASY(b *testing.B) {
	for i := 0; i < b.N; i++ {

		jso___EASY()
	}
}
