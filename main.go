package main

import (
	"fmt"
	"go/token"
	"io/ioutil"
	"log"
	"os"
)

var fset = token.NewFileSet()

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to pass the file name")
	}

	bs, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	methods := Generate(bs)
	for _, method := range methods {
		fmt.Printf("%+v\n", method)
	}
}

