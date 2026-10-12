package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func Process() error {
	inputYaml := "input.yaml"

	inputRaw, err := os.ReadFile(inputYaml)
	if err != nil {
		return fmt.Errorf("Fail to load input %s, %w", inputYaml, err)
	}
	var context map[string]any
	err = yaml.Unmarshal(inputRaw, &context)
	if err != nil {
		return fmt.Errorf("Fail to unmarshal input %s, %w", inputYaml, err)
	}

	fmt.Printf("%s\n", inputRaw)

	return nil

}
