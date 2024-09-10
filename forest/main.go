//go:generate easyjson -all main.go
package forest

import (
	"encoding/json"
	"github.com/mailru/easyjson"
)

type Forest struct {
	Name1 string
	Rio1  []string
	Name2 string
	Rio2  []string
	Name3 string
	Rio3  []string
	Name4 string
	Rio4  []string
	Name5 string
	Rio5  []string
}

func jso() {
	forest := Forest{Name1: "Egor"}
	marshal, err := json.Marshal(&forest)
	if err != nil {
		panic(err)
	}
	_ = marshal
}
func jso___EASY() {
	forest := Forest{Name1: "Egor"}
	marshal, err := easyjson.Marshal(&forest)
	if err != nil {
		panic(err)
	}
	_ = marshal
}
