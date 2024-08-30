package main

import (
	"log"
	"strings"
	"unicode/utf8"
)

func main() {

	//:::: make runa with builder
	var builder strings.Builder
	var runa []byte = make([]byte, 10)
	utf8.EncodeRune(runa, 2025)
	//:: write to builder
	builder.Write(runa)
	println(builder.String())
	//:: trim from builder '0' runa
	ru := strings.Trim(builder.String(), string(0))

	//	:::: NOW we have runa, which encode with not a signle byte

	nameWithRuna := "egor" + ru

	//	:::: ITERATE THROUGH ALL SYMBOLS

	//::; METH 1
	for _, r := range nameWithRuna {
		log.Println(string(r))
	}
	//::; METH 2
	slic := strings.Split(nameWithRuna, "")
	for _, s := range slic {
		log.Println(s)

	}
	//::; METH 3
	for len(nameWithRuna) > 0 {
		r, size := utf8.DecodeRuneInString(nameWithRuna)
		log.Printf("%c", r)
		nameWithRuna = nameWithRuna[size:]
	}

}

//println(strings.Index(name, "e"))

func PrintPointersOfelemntsSlice(rio *[]int) {
}
