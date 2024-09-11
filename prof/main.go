package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
)

var fileName = "gorik.txt"

func main() {

	runtime.SetCPUProfileRate(10000)
	file, _ := os.Create("cpu.pp")
	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()
	for i := 0; i < 10; i++ {

		lines := SomeFunc(fileName)
		fmt.Println(lines)
	}

}

func SomeFunc(filePath string) []string {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)

	}

	filesString := strings.Split(string(file), "\n")

	//bufferedWriter.Flush()
	//return writer

	return filesString
}
