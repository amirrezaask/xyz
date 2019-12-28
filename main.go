package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to pass the file name")
	}
	s, err := New("query", &funcTemplateData{
		SelfType:   "*db",
		Name:       "Insert",
		ReturnType: "[]Book",
		Query:      fmt.Sprintf("`%s`", selectGenerator("books", "FindByNameAndAge")),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	//bs, err := ioutil.ReadFile(os.Args[1])
	//if err != nil {
	//	log.Fatal(err)
	//}
	//methods := Generate(bs)
	//for _, method := range methods {
	//
	//}

}

var typ2typ = map[string]string{
	"Find":   "query",
	"Update": "exec",
	"Insert": "exec",
	"Delete": "exec",
}
