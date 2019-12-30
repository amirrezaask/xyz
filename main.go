package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var codes []string
	bs, _ := ioutil.ReadFile("book.go")
	methods, abstract, impl := Parse(bs)
	for _, method := range methods {
		templateData := &funcTemplateData{
			SelfType:   impl,
			Name:       method.name,
			ReturnType: method.returns,
			Query:      method.query,
		}
		fun, err := Generate(typ2typ[method.typ], templateData)
		if err != nil {
			panic(err)
		}
		codes = append(codes, fun)
	}
	sCode, err := Generate("struct", &funcTemplateData{
		SelfType: impl,
	})
	if err != nil {
		panic(err)
	}
	codes = append(codes, sCode)
	nCode, err := Generate("new", &newTemplateData{
		AbstractName: abstract,
		ImplName:     impl,
	})
	if err != nil {
		panic(err)
	}
	codes = append(codes, nCode)
	file, err := Generate("file", &fileTemplateData{
		PackageName: "db",
		Codes:       strings.Join(codes, "\n"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(file)
}

var typ2typ = map[string]string{
	"Find":   "query",
	"Update": "exec",
	"Insert": "exec",
	"Delete": "exec",
}
