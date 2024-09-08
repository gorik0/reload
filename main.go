package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"


func main() {
	log.Println(runtime.NumCPU())
}

func makeWordsFromFile_Scanenr(file *os.File) ([]string, error) {

	scanner := bufio.NewScanner(file)
	var words []string
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, nil

}
func makeWordsFromFile_Rural(file *os.File) ([]string, error) {
	var words []string
	r := bufio.NewReader(file)
	var word string
	for {
		_, err := fmt.Fscan(r, &word)

		if err != nil {
			if err == io.EOF {
				//words = append(words, word)

				return words, nil
			} else {
				return nil, err
			}

		}
		words = append(words, word)
	}

}
func makeWordsFromFile_String(file *os.File) ([]string, error) {

	writer := strings.Builder{}
	_, err := file.WriteTo(&writer)
	if err != nil {
		return nil, err
	}
	scanner := strings.Fields(writer.String())

	return scanner, nil

}
func makeWordsFromFile_Buffered(file *os.File) ([]string, error) {

	reader := bufio.NewReader(file)
	var words []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				words = append(words, strings.Split(line, " ")...)
				break
			} else {
				return nil, err
			}
		}
		words = append(words, strings.Split(line, " ")...)

	}
	return words, nil

}
