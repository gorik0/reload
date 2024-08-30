package main

import (
	"io"
	"log"
	"time"
)

func main() {

	//	::: MULTI reader
	//	::: MULTI writer
	r, w := io.Pipe()

	go writeToWriter(w)
	listeReader(r)

}

func listeReader(r *io.PipeReader) {
	for {
		all, err := io.ReadAll(r)
		if err != nil {
			log.Printf("Error reading from Reader: %v", err)
			return
		}
		log.Printf("From reader: %s", string(all))

	}

}

func writeToWriter(w *io.PipeWriter) {
	for {
		time.Sleep(time.Second)
		write, err := w.Write([]byte("Helo, goriko!!!"))
		if err != nil {
			log.Printf("error writing to pipe: %v", err)
			return
		}

		log.Printf("Has been written to pipe: %v", write)
	}

}
