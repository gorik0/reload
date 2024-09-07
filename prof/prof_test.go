package prof

import "testing"

func Benchmark___simple_allo(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			simple_allo(100000)
		}
	})
}
func Benchmark___Preallo(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Preallo(100000)
		}
	})
}
