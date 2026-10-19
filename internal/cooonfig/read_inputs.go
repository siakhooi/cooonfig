package cooonfig

import (
	"fmt"

	"path/filepath"
)

func isReserved(name string) bool {
	return name == model
}
func readInputs(inputs []input, inputDir string) (map[string]any, error) {

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

		if err := readYaml(inputYaml, &inputContext); err != nil {
			return nil, err
		}
		context[input.Name] = inputContext

	}
	return context, nil
}
