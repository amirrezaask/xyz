package main

import (
	"bytes"
	"html/template"
)

func New(templateName string, data interface{}) (string, error) {
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
	ReturnType string
}

type fileTemplateData struct {
}

const file = `package {{}}


`

const newFunc = `func New() ({{.ReturnType}}, error) {

}`
const execFunc = `func (r {{.SelfType}}) {{.Name}}(args ...interface{}) error {

}`
const queryFunc = `func (r {{.SelfType}}) {{.Name}}(args ...interface{}) ({{.ReturnType}}, error) {

}`
const repoStruct = `type {{.SelfType}} struct {
	db *sql.DB
}`

var Templates = map[string]string{
	"new":   newFunc,
	"exec":  execFunc,
	"query": queryFunc,
	"file":  file,
}
