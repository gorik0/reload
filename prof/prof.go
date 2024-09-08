package prof

import (
	"regexp"
	"strings"
	"sync"
)

var mutex = sync.Mutex{}

var compile *regexp.Regexp

func init() {
	compile, _ = regexp.Compile("egor")

}

func Reg_Compile(findOf, findIn string) {

	_ = compile.MatchString(findIn)

}
func Reg_NoCompile(findOf, findIn string) {
	_, _ = regexp.MatchString(findOf, findIn)

}
func No_Reg(findOf, findIn string) {
	_ = strings.Contains(findOf, findIn)

}
