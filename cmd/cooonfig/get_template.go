package main

import (
	"github.com/Masterminds/sprig/v3"
	"path/filepath"
	"text/template"
)

func GetTemplate(templateFilename string) (*template.Template, error) {
	return template.New(filepath.Base(templateFilename)).
		Funcs(sprig.TxtFuncMap()).
		ParseFiles(templateFilename)
}
