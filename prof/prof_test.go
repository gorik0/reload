package prof

import "testing"

func BenchmarkLOCK(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			WithLock()
		}
	})
}
func BenchmarkUNlock(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			With_NO_Lock()
		}
	})
}
