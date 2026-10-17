package main

import (
	"fmt"
	"os"
	"text/template"
)

func GenerateOutputs(model map[string]any, outputs []Output) error {

	outputTemplateYaml := "output.json.template"
	t1, err := template.ParseFiles(outputTemplateYaml)
	if err != nil {
		return fmt.Errorf("fail to load output template %s: %w", outputTemplateYaml, err)
	}
	outputFile := "output.json"
	file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("fail to open file %s: %w", outputFile, err)
	}
	err = t1.Execute(file, model)
	if err != nil {
		return fmt.Errorf("fail to render output template %s: %w", outputTemplateYaml, err)
	}

	return nil

}
