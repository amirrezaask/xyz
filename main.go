package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	pf, err := parser.ParseFile(fset, "", bs, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	//ast.Print(fset,pf)
	var xyzDecs []ast.Decl
	for d := range pf.Decls {
		if strings.Contains(pf.Decls[d].(*ast.GenDecl).Doc.Text(), "@xyz") {
			xyzDecs = append(xyzDecs, pf.Decls[d])
		}
	}
	var model *ast.StructType
	var repo *ast.InterfaceType
	var name string
	_ , _ = model, repo
	for x := range xyzDecs {
		typeSpec := xyzDecs[x].(*ast.GenDecl).Specs[0].(*ast.TypeSpec)
		asStruct, isStruct := typeSpec.Type.(*ast.StructType)
		if isStruct {
			name = typeSpec.Name.Name
			model = asStruct
			continue
		}
		asInterface, isInterface := typeSpec.Type.(*ast.InterfaceType)
		if isInterface {
			repo = asInterface
		}
	}
	methods := getListOfMethodsOfInterface(repo)
	methods = append(methods, "INSERT")
	fields := getListOfFields(model)
	for m := range methods {
		fmt.Println(generate(name, methods[m], fields))
	}
}

