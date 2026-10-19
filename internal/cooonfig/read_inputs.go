package cooonfig

import (
	"fmt"

	"path/filepath"
)

func readInputs(inputs []input, inputDir string) (map[string]any, error) {

	context := map[string]any{}
	for _, input := range inputs {

		inputYaml := filepath.Join(inputDir, input.Path)
		fmt.Printf("read input: %s: %s\n", input.Name, inputYaml)

		var inputContext map[string]any

		if err := readYaml(inputYaml, &inputContext); err != nil {
			return nil, err
		}
		context[input.Name] = inputContext

	}
	return context, nil
}
