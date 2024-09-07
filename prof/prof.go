package prof

import "sync"

var mutex = sync.Mutex{}

func Preallo(size int) []int {
	a := make([]int, 0, size)
	for i := 0; i < size; i++ {
		a = append(a, i)

	}
	return a
}
func simple_allo(size int) []int {
	a := make([]int, 0)
	for i := 0; i < size; i++ {
		a = append(a, i)

	}
	return a
}
