package prof

import "sync"

var mutex = sync.Mutex{}

func With_NO_Lock() int {
	a := 0
	for i := 0; i < 100000; i++ {
		a++

	}
	return a
}
func WithLock() int {
	mutex.Lock()
	defer mutex.Unlock()
	a := 0
	for i := 0; i < 100000; i++ {
		a++

	}
	return a
}
