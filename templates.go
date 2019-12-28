package main

import (
	"bytes"
	"html/template"
	"strings"
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
	Query      string
}

func (f *funcTemplateData) IsSlice(s string) bool {
	return strings.Contains(s, "[]")
}

type fileTemplateData struct {
}

const file = `package {{}}


`

const newFunc = `func New(db *sql.DB) ({{.ReturnType}}, error) {
	return &{{.SelfType}}{
		db
	}
}`
const execFunc = `func (r {{.SelfType}}) {{.Name}}(args ...interface{}) error {
	_, err := db.Exec({{.Query}})
	if err != nil {
		return err
	}
	return nil
}`
const queryFunc = `func (r {{.SelfType}}) {{.Name}}(args ...interface{}) ({{.ReturnType}}, error) {
	var ret {{.ReturnType}}
	{{if .IsSlice .ReturnType}}
	res, err := db.Query({{.Query}}, args...)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		m := new({{.ReturnType}})
		err = res.Scan(m)
		if err != nil {
			return nil, err
		}
		ret = append(ret, m)
	}
	return ret, nil
	{{else}}
	res, err := db.QueryRow({{.Query}}, args...)
	if err != nil {
		return nil, err
	}
	err = res.Scan(m)
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
	"new":   newFunc,
	"exec":  execFunc,
	"query": queryFunc,
	"file":  file,
}
