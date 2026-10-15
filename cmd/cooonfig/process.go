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
func Process(configFile string) error {
	fmt.Printf("configFile: %s\n", configFile)

	cfg, err := ReadConfig(configFile)
	if err != nil {
		return fmt.Errorf("fail to read Config %s: %w", configFile, err)
	}
	fmt.Printf("test read: %s:%s\n", cfg.Inputs[0].Name, cfg.Inputs[0].Path)

	inputYaml := "input.yaml"

	inputRaw, err := os.ReadFile(inputYaml)
	if err != nil {
		return fmt.Errorf("fail to load input %s: %w", inputYaml, err)
	}
	var context map[string]any
	err = yaml.Unmarshal(inputRaw, &context)
	if err != nil {
		return fmt.Errorf("fail to unmarshal input %s: %w", inputYaml, err)
	}
	context["final"] = map[string]any{}

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
	context["final"] = merge(context["final"].(map[string]any), model)

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

	fmt.Printf("model: %v\n", context["final"])
	err = t1.Execute(file, context["final"])
	if err != nil {
		return fmt.Errorf("fail to render output template %s: %w", outputTemplateYaml, err)
	}

	return nil

}
