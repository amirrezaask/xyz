package main

import (
	"bytes"
	"strings"
	"text/template"
)

func Generate(templateName string, data interface{}) (string, error) {
	t, err := template.New("tpl").Parse(Templates[templateName])
	if err != nil {
		return "", err
	}
	var w bytes.Buffer
	err = t.Execute(&w, data)
	if err != nil {
		return "", err
	}
	return w.String(), nil
}

type funcTemplateData struct {
	SelfType   string
	Name       string
	ReturnType []string
	Query      string
}

func (f *funcTemplateData) ReturnWithoutBracket() string {
	return f.ReturnType[0][2:]
}
func (f *funcTemplateData) ReturnsCommaSeperated() string {
	return strings.Join(f.ReturnType, ", ")
}
func (f *funcTemplateData) ReturnWithoutError() string {
	return f.ReturnType[0]
}
func (f *funcTemplateData) IsSlice(s []string) bool {
	if len(s) > 2 {
		panic("length of return types should not be greater than 2")
	}
	return strings.Contains(s[0], "[]")
}

type fileTemplateData struct {
	PackageName string
	Codes       string
}

const file = `package {{.PackageName}}
	{{.Codes}}

`

type newTemplateData struct {
	AbstractName, ImplName string
}

const newFunc = `func New{{.AbstractName}}(db *sql.DB) ({{.AbstractName}}, error) {
	return &{{.ImplName}}{
		db
	}
}`
const execFunc = `func (r *{{.SelfType}}) {{.Name}}(args ...interface{}) error {
	_, err := db.Exec("{{.Query}}")
	if err != nil {
		return err
	}
	return nil
}`
const queryFunc = `func (r *{{.SelfType}}) {{.Name}}(args ...interface{}) ({{.ReturnsCommaSeperated}}) {
	var ret {{.ReturnWithoutError}}
	{{if .IsSlice .ReturnType}}
	err := db.Select(&ret, "{{.Query}}", args...)
	if err != nil {
		return nil, err
	}
	return ret, nil
	{{else}}
	res, err := db.Get(&ret, "{{.Query}}", args...)
	if err != nil {
		return nil, err
	}
	return ret, nil
	{{end}}
}`

const repoStruct = `type {{.SelfType}} struct {
	db *sql.DB
}`

var Templates = map[string]string{
	"new":    newFunc,
	"exec":   execFunc,
	"query":  queryFunc,
	"file":   file,
	"struct": repoStruct,
}
