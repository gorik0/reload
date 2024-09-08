package prof

import (
	"bytes"
	"regexp"
	"strings"
	"sync"
)

var mutex = sync.Mutex{}

var compile *regexp.Regexp

func init() {
	compile, _ = regexp.Compile("egor")

}

func SimpleConcat() string {
	return "3" + "3" + "3"

}
func ThroughBuilder() string {
	builder := strings.Builder{}
	builder.WriteString("3")
	builder.WriteString("3")
	builder.WriteString("3")
	return builder.String()

}

func ThroughBuffer() string {

	buffer := bytes.Buffer{}
	buffer.Grow(10)
	buffer.WriteString("3")
	buffer.WriteString("3")
	buffer.WriteString("3")
	return buffer.String()
}
func Assignment() string {
	a := ""
	a += "3"
	a += "3"
	a += "3"
	return a
}
