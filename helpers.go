package main

import (
	"go/ast"
	"regexp"
	"strings"
)

func getListOfMethodsOfInterface(i *ast.InterfaceType) []string {
	var names []string
	for l := range i.Methods.List {
		names = append(names, i.Methods.List[l].Names[0].Name)
	}
	return names
}

func getListOfFields(m *ast.StructType) []string {
	var fields []string
	for f := range m.Fields.List {
		names :=  m.Fields.List[f].Names
		for _, name := range names {
			fields = append(fields, ToSnakeCase(name.Name))
		}
	}
	return fields
}
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake  = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}