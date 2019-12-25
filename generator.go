package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"log"
	"strings"
)

const SELECT = "SELECT * FROM %s WHERE %s"
const DELETE = "DELETE FROM %s WHERE %s"
const UPDATE = "UPDATE %s SET "
const INSERT = "INSERT INTO %s (%s) VALUES (%s)"
type method struct {
	name, query string
}
func Generate(bs []byte) []*method {
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
	var generatedMethods []*method
	for m := range methods {
		generatedMethods = append(generatedMethods, &method{methods[m], generate(name, methods[m], fields)})
	}
	return generatedMethods
}
func generate(tableName, name string, fields []string) string {
	if name[:4] == "Find" {
		return selectGenerator(tableName, name)
	} else if name[:6] == "Update" {
		return updateGenerator(tableName, name)
	} else if name[:6] == "Delete" {
		return deleteGenerator(tableName, name)
	} else if name[:6] == "INSERT" {
		return insertGenerator(tableName, fields)
	} else {
		log.Fatalf("Method name %s is not valid\n", name)
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
	for _, updateParam := range updateParams{
		updateParamslist = append(updateParamslist, fmt.Sprintf("%s = ?", updateParam))
	}

	query += strings.Join(updateParamslist, ", ") + " WHERE "

	var selectParamslist []string
	for _, selectParam := range selectParams{
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
	return fmt.Sprintf(INSERT, tableName, strings.Join(fields, ", "),strings.Join(questions, ", "))
}
