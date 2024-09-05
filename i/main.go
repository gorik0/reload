package main

import (
	"os"
)

func main() {

	file, err := os.OpenFile("be.txt", os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 1000000; i++ {

		_, err := file.Write([]byte("helo "))
		if err != nil {
			panic(err)
		}
	}

}
