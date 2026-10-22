package cooonfig

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func templateFunctionRequired(message string, value any) (any, error) {
	if value == nil || value == "" {
		return nil, fmt.Errorf("%s", message)
	}
	return value, nil
}

func templateFunctionToYaml(value any) (string, error) {
	data, err := yaml.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
