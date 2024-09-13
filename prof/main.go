package main

import (
	"math/rand"
	_ "net/http/pprof"
	"sync"
)

var (
	size    = 100000
	wg      sync.WaitGroup
	goruNum = 100
)

func main() {
	rio := make([]int, size)
	wg.Add(goruNum)

	for i := 0; i < goruNum; i++ {
		go RioChangeMake(rio)
	}
}

func RioChangeMake(rio []int) {
	var i int = 0
	for i < size {
		rio[i] = rand.Intn(size)

		i += rand.Intn(10)

	}
}
