package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func GenerateOutputs(model map[string]any, outputs []Output) error {
	for _, output := range outputs {

		fmt.Printf("render output: %s => %s\n", output.Template, output.OutFile)

		templateFile := output.Template

		t1, err := template.ParseFiles(templateFile)
		if err != nil {
			return fmt.Errorf("fail to load output template %s: %w", templateFile, err)
		}
		outputFile := output.OutFile
		dir := filepath.Dir(outputFile)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %q: %w", dir, err)
		}
		file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("fail to open file %s: %w", outputFile, err)
		}
		err = t1.Execute(file, model)
		if err != nil {
			return fmt.Errorf("fail to render output template %s: %w", templateFile, err)
		}

	}
	return nil

}
