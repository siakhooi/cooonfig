package main

import (
	"text/template"
)

func GetTemplate(templateFilename string) (*template.Template, error) {
	return template.ParseFiles(templateFilename)
}
