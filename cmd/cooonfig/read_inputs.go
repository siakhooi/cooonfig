package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func ReadInputs(inputs []Input, inputDir string) (map[string]any, error) {

	var context = map[string]any{}
	context[MODEL] = map[string]any{}

	for _, input := range inputs {

		inputYaml := filepath.Join(inputDir, input.Path)
		fmt.Printf("read input: %s: %s\n", input.Name, inputYaml)

		if context[input.Name] != nil {
			return nil, fmt.Errorf("input name already exists or is a reserved word: %s", input.Name)
		}
		var inputContext map[string]any

		inputRaw, err := os.ReadFile(inputYaml)
		if err != nil {
			return nil, fmt.Errorf("read file %q: %w", inputYaml, err)
		}
		err = yaml.Unmarshal(inputRaw, &inputContext)
		if err != nil {
			return nil, fmt.Errorf("unmarshal %q: %w", inputYaml, err)
		}
		context[input.Name] = inputContext

	}
	return context, nil
}
