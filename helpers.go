package main

import (
	"go/ast"
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
			fields = append(fields, name.Name)
		}
	}
	return fields
}