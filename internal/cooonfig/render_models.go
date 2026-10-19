package cooonfig

import (
	"bytes"
	"fmt"
	"path/filepath"
)

const model = "model"

func renderModel(context map[string]any, modelFiles []string, modelDir string) (map[string]any, error) {
	modelMap := map[string]any{}

	for _, modelTemplate := range modelFiles {
		modelTemplatePath := filepath.Join(modelDir, modelTemplate)

		fmt.Printf("read model: %s\n", modelTemplatePath)

		t, err := getTemplate(modelTemplatePath)
		if err != nil {
			return nil, fmt.Errorf("get template %q: %w", modelTemplate, err)
		}

		context[model] = modelMap
		var model map[string]any
		var buf bytes.Buffer
		err = t.Execute(&buf, context)
		if err != nil {
			return nil, fmt.Errorf("execute template %q: %w", modelTemplate, err)
		}
		if err := unmarshalYaml(modelTemplate, buf.Bytes(), &model); err != nil {
			return nil, err
		}
		deepMerge(modelMap, model)
	}
	return modelMap, nil

}
