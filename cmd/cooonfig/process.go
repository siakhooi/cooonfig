package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"text/template"
)

func Process(configFile string) error {

	cfg, err := ReadConfig(configFile)
	if err != nil {
		return fmt.Errorf("fail to read Config %s: %w", configFile, err)
	}
	context, err := ReadInputs(cfg.Inputs)
	if err != nil {
		return fmt.Errorf("fail to read input files: %w", err)
	}

	models, err := RenderModels(context, cfg.Models)
	if err != nil {
		return fmt.Errorf("fail to render models: %w", err)
	}

	outputTemplateYaml := "output.json.template"
	t1, err1 := template.ParseFiles(outputTemplateYaml)
	if err1 != nil {
		return fmt.Errorf("fail to load output template %s: %w", outputTemplateYaml, err1)
	}
	outputFile := "output.json"
	file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("fail to open file %s: %w", outputFile, err)
	}
	b, err := yaml.Marshal(models)
	fmt.Printf("model:\n%v\n", string(b))
	err = t1.Execute(file, models)
	if err != nil {
		return fmt.Errorf("fail to render output template %s: %w", outputTemplateYaml, err)
	}

	return nil

}
