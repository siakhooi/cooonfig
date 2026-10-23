package cooonfig

import (
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func buildTemplateFunctionMap() template.FuncMap {
	funcMap := sprig.TxtFuncMap()

	funcMap["required"] = templateFunctionRequired
	funcMap["toYaml"] = templateFunctionToYaml

	return funcMap
}
func getTemplate(templateFilename string) (*template.Template, error) {
	funcMap := buildTemplateFunctionMap()

	return template.New(filepath.Base(templateFilename)).
		Funcs(funcMap).
		ParseFiles(templateFilename)
}
