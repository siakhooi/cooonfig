package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"maps"
	"os"
	"text/template"
)

func merge(m1, m2 map[string]any) map[string]any {
	return maps.Clone(m2)
}

const MODEL = "model"

func Process(configFile string) error {

	cfg, err := ReadConfig(configFile)
	if err != nil {
		return fmt.Errorf("fail to read Config %s: %w", configFile, err)
	}
	context, err := ReadInputs(cfg.Inputs)
	if err != nil {
		return fmt.Errorf("fail to read input files: %w", err)
	}

	templateYaml := "template.yaml"
	t, err := template.ParseFiles(templateYaml)
	if err != nil {
		return fmt.Errorf("fail to load template %s: %w", templateYaml, err)
	}

	var model map[string]any
	var buf bytes.Buffer
	err = t.Execute(&buf, context)
	if err != nil {
		return fmt.Errorf("fail to render template %s: %w", templateYaml, err)
	}
	if err := yaml.Unmarshal(buf.Bytes(), &model); err != nil {
		return fmt.Errorf("fail to convert to model %s: %w", templateYaml, err)
	}
	context[MODEL] = merge(context[MODEL].(map[string]any), model)

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
	b, err := yaml.Marshal(context[MODEL])
	fmt.Printf("model:\n%v\n", string(b))
	err = t1.Execute(file, context[MODEL])
	if err != nil {
		return fmt.Errorf("fail to render output template %s: %w", outputTemplateYaml, err)
	}

	return nil

}
