package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to pass the file name")
	}
	//s, err := New("exec", funcTemplateData{
	//	SelfName:   "d",
	//	SelfType:   "*db",
	//	Name:       "Insert",
	//	ReturnType: "",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(s)
	bs, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	methods := Generate(bs)
	for _, method := range methods {
		fmt.Println(typ2typ[method.typ])
	}

}

var typ2typ = map[string]string{
	"Find":   "query",
	"Update": "exec",
	"Insert": "exec",
	"Delete": "exec",
}
