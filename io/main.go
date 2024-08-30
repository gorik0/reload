package main

import (
	"io"
	"os"
	"strings"
)

func main() {

	//	:::: string BUIDLER make concat less memory
	var bi strings.Builder

	bi.Write([]byte("egor"))

	io.Copy(os.Stdout, strings.NewReader(bi.String()))
}

//println(strings.Index(name, "e"))

func PrintPointersOfelemntsSlice(rio *[]int) {
}
