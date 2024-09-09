package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {
	printAlloc()
	for i := 0; i < 10; i++ {

		makeRIOandCOUNT(1e4)

	}
	printAlloc()

}

//go:noinline
func makeRIOandCOUNT(num int) {
	printAlloc()
	fmt.Println("BEFORE")

	var res int
	rio := make([]int, num)
	for i := range len(rio) {
		rio[i] = rand.Intn(num)

	}

	for _, i := range rio {
		res += i

	}
	_ = res

	printAlloc()
	fmt.Println("AFTER")
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}

var pool = sync.Pool{New: func() interface{} {
	fmt.Println("NEW again")
	rio := make([]int, 1e4)
	return rio
}}

//go:noinline
func makeRIOandCOUNT_____sync(num int) {
	printAlloc()
	fmt.Println("BEFORE")
	var res int
	rio := pool.Get().([]int)
	defer pool.Put(rio)
	for i := range len(rio) {
		rio[i] = rand.Intn(num)

	}

	for _, i := range rio {
		res += i

	}
	_ = res
	printAlloc()
	fmt.Println("AFTER")
}
