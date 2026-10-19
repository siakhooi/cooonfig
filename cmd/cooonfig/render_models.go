package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"path/filepath"
)

const MODEL = "model"

func RenderModel(context map[string]any, modelFiles []string, modelDir string) (map[string]any, error) {
	for _, modelTemplate := range modelFiles {
		modelTemplatePath := filepath.Join(modelDir, modelTemplate)

		fmt.Printf("read model: %s\n", modelTemplatePath)

		t, err := GetTemplate(modelTemplatePath)
		if err != nil {
			return nil, fmt.Errorf("get template %q: %w", modelTemplate, err)
		}

		var model map[string]any
		var buf bytes.Buffer
		err = t.Execute(&buf, context)
		if err != nil {
			return nil, fmt.Errorf("execute template %q: %w", modelTemplate, err)
		}
		if err := yaml.Unmarshal(buf.Bytes(), &model); err != nil {
			return nil, fmt.Errorf("unmarshal %q: %w", modelTemplate, err)
		}
		DeepMerge(context[MODEL].(map[string]any), model)
	}
	return context[MODEL].(map[string]any), nil

}
