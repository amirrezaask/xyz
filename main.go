package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var codes []string
	if len(os.Args) < 2 {
		log.Fatal("Usage:\n xyz filename.go")
	}
	bs, _ := ioutil.ReadFile(os.Args[1])
	methods, abstract, impl, pkg := Parse(bs)
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
		PackageName: pkg,
		Codes:       strings.Join(codes, "\n"),
	})
	if err != nil {
		panic(err)
	}
	generatedFileName := strings.Split(os.Args[1], ".")[0] + "_xyz.go"
	err = ioutil.WriteFile(generatedFileName, []byte(file), 0644)
	if err != nil {
		panic(err)
	}

}

var typ2typ = map[string]string{
	"Find":   "query",
	"Update": "exec",
	"Insert": "exec",
	"Delete": "exec",
}
