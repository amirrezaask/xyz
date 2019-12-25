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
	ast.Print(fset,pf)
	var xyzDecs []ast.Decl
	for d := range pf.Decls {
		if strings.Contains(pf.Decls[d].(*ast.GenDecl).Doc.Text(), "@xyz") {
			xyzDecs = append(xyzDecs, pf.Decls[d])
		}
	}
	var model *ast.StructType
	var repo *ast.InterfaceType
	_ , _ = model, repo
	for x := range xyzDecs {
		typeSpec := xyzDecs[x].(*ast.GenDecl).Specs[0].(*ast.TypeSpec)
		asStruct, isStruct := typeSpec.Type.(*ast.StructType)
		if isStruct {
			model = asStruct
		}
		asInterface, isInterface := typeSpec.Type.(*ast.InterfaceType)
		if isInterface {
			repo = asInterface
		}
	}
	methods := getListOfMethodsOfInterface(repo)
}

const SELECT = "SELECT * FROM %s WHERE %s"
func selectGenerator(tableName, method string) string {
	queryParams := strings.Split(method, "By")[1]
	params := strings.Split(queryParams, "And")
	var paramsAndQuestions []string
	for p := range params {
		paramsAndQuestions = append(paramsAndQuestions, fmt.Sprintf("%s=?", params[p]))
	}
	return fmt.Sprintf(SELECT, tableName, strings.Join(paramsAndQuestions, "AND"))
}
const DELETE = "DELETE FROM %s WHERE %s"
func deleteGenerator(tableName, method string) string {
	queryParams := strings.Split(method, "By")[1]
	params := strings.Split(queryParams, "And")
	var paramsAndQuestions []string
	for p := range params {
		paramsAndQuestions = append(paramsAndQuestions, fmt.Sprintf("%s=?", params[p]))
	}
	return fmt.Sprintf(DELETE, tableName, strings.Join(paramsAndQuestions, "AND"))
}
func parseUpdateMethod(method string) []string {
	//UpdateNameAndFNameBasedOnAge
	selectParams := strings.Split(strings.Split(method, "BasedOn")[1], "And")
	updateParams := (strings.Split(method, "BasedOn")[0])[6:]

}
type query struct {
	typ string
	params []string
}
func parseMethodName(methods []string) query{
	//find
	//update
	//delete
	for _, name := range methods {
		if name[:4] == "Find" {
			return query{"select",parseFindMethod(name)}
		} else if name[:6] == "Update" {
			return query{
				typ:    "",
				params: nil,
			}
		} else if name[:6] == "Delete" {

		} else {
			log.Fatalf("Method name %s is not valid\n", name)
		}
	}
}

func getListOfMethodsOfInterface(i *ast.InterfaceType) []string {
	var names []string
	for l := range i.Methods.List {
		names = append(names, i.Methods.List[l].Names[0].Name)
	}
	return names
}
