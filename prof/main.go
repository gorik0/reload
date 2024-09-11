package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var fileName = "gorik.txt"

func main() {
	//file, _ := os.Create("gorik.txt")
	//for i := 0; i < 100000; i++ {
	//	file.WriteString("helo gorik")
	//
	//	if i%10 == 0 {
	//		file.WriteString("\n")
	//	}
	//
	//}

	lines := SplitByLines___Split(fmt.Sprintf("gorik.txt"))
	fmt.Println(len(lines))
}

func SplitByLines___Buffer__BUFIO(filePath string) []string {

	var lines []string = make([]string, 0, 100)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)

	}

	bufferedReader := bufio.NewReader(file)

	//writer := strings.Builder{}
	//bufferedWriter := bufio.NewWriter(&writer)
	for {
		line, err := bufferedReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				//bufferedWriter.Write(line)
				lines = append(lines, line)
				break
			} else {
				panic(err)
			}
		}
		//bufferedWriter.Write(line)

		lines = append(lines, line)
	}
	//bufferedWriter.Flush()
	//return writer

	return lines
}

func SplitByLines___Buffer__BYTES(filePath string) []byte {

	var bytes = make([]byte, 0, 100)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)

	}

	bufferedReader := bufio.NewReader(file)

	//writer := strings.Builder{}
	//bufferedWriter := bufio.NewWriter(&writer)
	var byt = make([]byte, 0, 100)
	for {
		byt, err = bufferedReader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				//bufferedWriter.Write(line)
				bytes = append(bytes, byt...)
				break
			} else {
				panic(err)
			}
		}
		//bufferedWriter.Write(line)

		bytes = append(bytes, byt...)
	}
	//bufferedWriter.Flush()
	//return writer

	return bytes
}
func SplitByLines___Split(filePath string) []string {

	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)

	}

	filesString := strings.Split(string(file), "\n")

	//bufferedWriter.Flush()
	//return writer

	return filesString
}
func SplitByLines___StringScanner(filePath string) []string {
	var lines []string = make([]string, 15)
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)

	}
	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	i := 0
	for scanner.Scan() {
		lines[i] = scanner.Text()
		i++
	}

	//bufferedWriter.Flush()
	//return writer

	return lines
}
