package main

import (
	"github.com/Masterminds/sprig/v3"
	"text/template"
)

func GetTemplate(templateFilename string) (*template.Template, error) {
	return template.New(templateFilename).
		Funcs(sprig.TxtFuncMap()).
		ParseFiles(templateFilename)
}
