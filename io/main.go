package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type egor int

func (e egor) Write(data []byte) (n int, err error) {
	log.Println(data)
	return len(data), nil
}
func main() {
	//::: bufio provide Buffered writer/reader to reduce memory usage , amount of sys call

	//	uth it's ntihing just a simple []byte, when it's overflowed it write to indoor writer

	buw := bufio.NewWriterSize(os.Stdout, 3)
	buw.WriteString("sdHEL")
	bi := strings.Builder{}

	//:::; reset this buffer, free buffer size and change inner writer (free data is lost, if not flushed)
	buw.Reset(&bi)

	buw.WriteString("and egor")

	print(bi.String())
	//io.WriteString(bufferedWriter, "egoo")

}
