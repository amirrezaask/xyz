package main

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"
)

func getMethodsFromInterface(i *ast.InterfaceType, fields []string, modelName string) []*methodGenerator {
	var methods []*methodGenerator
	for l := range i.Methods.List {
		var params []string
		var results []string
		for _, p := range i.Methods.List[l].Type.(*ast.FuncType).Params.List {
			for _, n := range p.Names {
				params = append(params, n.Name)
			}
		}
		if len(i.Methods.List[l].Type.(*ast.FuncType).Results.List) > 0 {
			for _, r := range i.Methods.List[l].Type.(*ast.FuncType).Results.List {
				results = append(results, fmt.Sprint(r.Type))
			}
		}

		methods = append(methods, &methodGenerator{
			name:    i.Methods.List[l].Names[0].Name,
			typ:     typeOfMethod(i.Methods.List[l].Names[0].Name),
			args:    params,
			query:   generate(modelName, i.Methods.List[l].Names[0].Name, fields),
			fields:  fields,
			returns: results,
		})
	}
	return methods
}

func getListOfFields(m *ast.StructType) []string {
	var fields []string
	for f := range m.Fields.List {
		names := m.Fields.List[f].Names
		for _, name := range names {
			fields = append(fields, toSnakeCase(name.Name))
		}
	}
	return fields
}

func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
