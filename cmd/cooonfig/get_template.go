package main

import (
	"github.com/Masterminds/sprig/v3"
	"path/filepath"
	"text/template"
)

func GetTemplate(templateFilename string) (*template.Template, error) {
	funcMap := sprig.TxtFuncMap()

	funcMap["required"] = templateFunctionRequired

	return template.New(filepath.Base(templateFilename)).
		Funcs(funcMap).
		ParseFiles(templateFilename)
}
