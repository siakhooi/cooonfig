package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
)

const MODEL = "model"

func RenderModel(context map[string]any, modelFiles []string) (map[string]any, error) {
	for _, modelTemplate := range modelFiles {
		fmt.Printf("read model: %s\n", modelTemplate)

		t, err := GetTemplate(modelTemplate)
		if err != nil {
			return nil, fmt.Errorf("fail to parse template %s: %w", modelTemplate, err)
		}

		var model map[string]any
		var buf bytes.Buffer
		err = t.Execute(&buf, context)
		if err != nil {
			return nil, fmt.Errorf("fail to render template %s: %w", modelTemplate, err)
		}
		if err := yaml.Unmarshal(buf.Bytes(), &model); err != nil {
			return nil, fmt.Errorf("fail to convert to model %s: %w", modelTemplate, err)
		}
		DeepMerge(context[MODEL].(map[string]any), model)
	}
	return context[MODEL].(map[string]any), nil

}
