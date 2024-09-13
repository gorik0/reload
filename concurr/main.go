package main

import (
	"runtime"
	"sync"
	"time"
)

var (
	mu sync.RWMutex
	wg sync.WaitGroup
)

const numGoru = 10000

func main() {

	mi(10)

	wg.Wait()
}

func mi(msxPro int) {
	runtime.GOMAXPROCS(msxPro)
	rio := make([]int, 1000)
	wg.Add(numGoru * 2)
	for i := 0; i < numGoru; i++ {

		go func(j int) {

			defer wg.Done()

			mu.Lock()
			SecretDo(j, rio)
			mu.Unlock()

		}(i)
	}
	for i := 0; i < numGoru; i++ {
		go func() {
			time.Sleep(time.Millisecond)
			defer wg.Done()

			mu.RLock()
			notSecretDo(rio)
			mu.RUnlock()

		}()
	}
}

func notSecretDo(rio []int) {

}

func SecretDo(j int, rio []int) {

	time.Sleep(time.Millisecond)
	for i := range rio {
		rio[i] = j
	}

}
