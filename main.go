package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, _ := ioutil.ReadFile("book.go")
	methods := Generate(bs)
	for _, method := range methods {
		fmt.Printf("%+v\n", method)
	}
}

var typ2typ = map[string]string{
	"Find":   "query",
	"Update": "exec",
	"Insert": "exec",
	"Delete": "exec",
}
