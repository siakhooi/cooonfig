package cooonfig

import (
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func getTemplate(templateFilename string) (*template.Template, error) {
	funcMap := sprig.TxtFuncMap()

	funcMap["required"] = templateFunctionRequired

	return template.New(filepath.Base(templateFilename)).
		Funcs(funcMap).
		ParseFiles(templateFilename)
}
