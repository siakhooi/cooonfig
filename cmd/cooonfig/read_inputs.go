package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func isReserved(name string) bool {
	return name == MODEL
}
func ReadInputs(inputs []Input, inputDir string) (map[string]any, error) {

	context := map[string]any{}
	for _, input := range inputs {

		inputYaml := filepath.Join(inputDir, input.Path)
		fmt.Printf("read input: %s: %s\n", input.Name, inputYaml)

		if isReserved(input.Name) {
			return nil, fmt.Errorf("input name is a reserved: %s", input.Name)
		}

		if context[input.Name] != nil {
			return nil, fmt.Errorf("input name already exists: %s", input.Name)
		}
		if input.Type != inputTypeYaml {
			return nil, fmt.Errorf("invalid type %s: %s", input.Name, input.Type)
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
