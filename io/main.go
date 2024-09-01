package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
)

type egor int

func (e egor) Read(p []byte) (n int, err error) {
	log.Println("read invoked")
	p[0] = 99
	switch {
	case len(p) > 10:
		return len(p), io.EOF
	default:
		return len(p), nil
	}
}

func main() {
	//::: bufio provide Buffered writer/reader to reduce memory usage , amount of sys call

	//	uth it's ntihing just a simple []byte, when it's overflowed it write to indoor writer
	//var reader egor
	//var strR = strings.NewReader("egor")
	//::: min size  16 byte
	//::: when request to reader first, invoke Read from internal reeader, then again, if buffer is empty

	//log.Println(bur.WriteTo(os.Stdout))
	//bur.Discard(21)
	reader := bytes.NewReader([]byte{112, 2, 31, 2, 23, 12, 12, 233})
	bur := bufio.NewReaderSize(reader, 3)
	//::: at most one call to internal reader

	line, err := bur.ReadBytes(2)
	if err != nil {
		log.Fatal(err)
	}
	i := make([]byte, 1)
	bur.Read(i)
	log.Println(i)
	line2, err := bur.ReadBytes(2)
	bur.Discard(5)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("line ::: ", line)
	log.Println("line2 ::: ", line2)
	//bur.WriteTo(os.Stdout)
	//println([]byte(string('n')))
}
