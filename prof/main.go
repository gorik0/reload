package main

import (
	"log"
	"math/rand"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/trace"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	debug.SetGCPercent(10)
	runtime.GOMAXPROCS(30)
	size := 10

	file, _ := os.Create("trace.ou")

	trace.Start(file)
	defer trace.Stop()
	wg.Add(size)
	for i := 0; i < size; i++ {
		go Goriko(10000)

	}

	log.Println("be_ FINISH")
	wg.Wait()
	log.Println("FINISH")
	defer runtime.StopTrace()

	println("done")
}

func Goriko(size int) {
	rio := make([]int, size)

	for i := range rio {

		rio[i] = rand.Intn(size)
	}
	println("BEFORE")
	//log.Println(rio)

	tick := time.Tick(time.Second * time.Duration(rand.Intn(10)))
	timer := time.NewTimer(time.Second * 10)

	makeSort(rio)
	println("AFTER")
	//log.Println(rio)
gorik:
	for {

		select {
		case <-tick:
			{
				insideTICK()
			}
		case <-timer.C:
			println("STOP")
			break gorik

		default:

		}
	}
	wg.Done()
}

func insideTICK() {
	time.Sleep(time.Second * 3)
}

func makeSort(rio []int) {
	for range rio {
		for i := range len(rio) - 1 {
			if rio[i] < rio[i+1] {
				rio[i], rio[i+1] = rio[i+1], rio[i]
			}

		}
	}

}
