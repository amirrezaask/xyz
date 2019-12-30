package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, _ := ioutil.ReadFile("book.go")
	methods := Parse(bs)
	for _, method := range methods {
		templateData := &funcTemplateData{
			SelfType:   method.selfType,
			Name:       method.name,
			ReturnType: method.returns,
			Query:      method.query,
		}
		fun, err := Generate(typ2typ[method.typ], templateData)
		if err != nil {
			panic(err)
		}
		fmt.Println(fun)
	}

}

var typ2typ = map[string]string{
	"Find":   "query",
	"Update": "exec",
	"Insert": "exec",
	"Delete": "exec",
}
