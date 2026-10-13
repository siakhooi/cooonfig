package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"text/template"
)

func merge(m1, m2 map[string]any) {
	m1 = m2
}
func Process() error {
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
		return fmt.Errorf("fail to template %s: %w", templateYaml, err)
	}
	if err := yaml.Unmarshal(buf.Bytes(), &model); err != nil {
		return fmt.Errorf("fail to convert to model %s: %w", templateYaml, err)
	}
	merge(context["final"].(map[string]any), model)

	fmt.Printf("%v\n", context)
	fmt.Printf("%v\n", model)

	return nil

}
