package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

const SELECT = "SELECT * FROM %s WHERE %s"
const DELETE = "DELETE FROM %s WHERE %s"
const UPDATE = "UPDATE %s SET "
const INSERT = "INSERT INTO %s (%s) VALUES (%s)"

type methodGenerator struct {
	name, typ string
	args      []string
	fields    []string
	returns   []string
	query     string
}

var fset = token.NewFileSet()

func Generate(bs []byte) []*methodGenerator {
	pf, err := parser.ParseFile(fset, "", bs, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	var xyzDecs []ast.Decl
	for d := range pf.Decls {
		if strings.Contains(pf.Decls[d].(*ast.GenDecl).Doc.Text(), "@xyz") {
			xyzDecs = append(xyzDecs, pf.Decls[d])
		}
	}
	var model *ast.StructType
	var repo *ast.InterfaceType
	var name string
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
	//			query: generate(name, methods[m], fields),

	fields := getListOfFields(model)
	methods := getMethodsFromInterface(repo, fields, name)

	return methods

}
func generate(tableName, methodName string, fields []string) string {
	typ := typeOfMethod(methodName)
	if typ == "Find" {
		return selectGenerator(tableName, methodName)
	} else if typ == "Update" {
		return updateGenerator(tableName, methodName)
	} else if typ == "Delete" {
		return deleteGenerator(tableName, methodName)
	} else if typ == "Insert" {
		return insertGenerator(tableName, fields)
	} else {
		log.Fatalf("Method methodName %s is not valid\n", typ)
		return ""
	}
}
func selectGenerator(tableName, method string) string {
	queryParams := strings.Split(method, "By")[1]
	params := strings.Split(queryParams, "And")
	var paramsAndQuestions []string
	for p := range params {
		paramsAndQuestions = append(paramsAndQuestions, fmt.Sprintf("%s=?", params[p]))
	}
	return fmt.Sprintf(SELECT, tableName, strings.Join(paramsAndQuestions, " AND "))
}
func deleteGenerator(tableName, method string) string {
	queryParams := strings.Split(method, "By")[1]
	params := strings.Split(queryParams, "And")
	var paramsAndQuestions []string
	for p := range params {
		paramsAndQuestions = append(paramsAndQuestions, fmt.Sprintf("%s=?", params[p]))
	}
	return fmt.Sprintf(DELETE, tableName, strings.Join(paramsAndQuestions, "AND"))
}
func updateGenerator(tableName string, method string) string {
	selectParams := strings.Split(strings.Split(method, "BasedOn")[1], "And")
	updateParams := strings.Split(strings.Split(method, "BasedOn")[0][6:], "And")

	query := fmt.Sprintf(UPDATE, tableName)

	var updateParamslist []string
	for _, updateParam := range updateParams {
		updateParamslist = append(updateParamslist, fmt.Sprintf("%s = ?", updateParam))
	}

	query += strings.Join(updateParamslist, ", ") + " WHERE "

	var selectParamslist []string
	for _, selectParam := range selectParams {
		selectParamslist = append(selectParamslist, fmt.Sprintf("%s = ?", selectParam))
	}
	query += strings.Join(selectParamslist, " AND ")

	return query
}
func insertGenerator(tableName string, fields []string) string {
	var questions []string
	for range fields {
		questions = append(questions, "?")
	}
	return fmt.Sprintf(INSERT, tableName, strings.Join(fields, ", "), strings.Join(questions, ", "))
}
func typeOfMethod(name string) string {
	if name[:4] == "Find" {
		return "Find"
	} else if name[:6] == "Update" {
		return "Update"
	} else if name[:6] == "Delete" {
		return "Delete"
	} else if name[:6] == "Insert" {
		return "Insert"
	} else {
		return ""
	}
}
