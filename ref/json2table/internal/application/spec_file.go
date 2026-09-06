/*
Package application run the application
*/
package application

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
)

func validateSpec(spec *Spec) error {
	if spec == nil {
		return fmt.Errorf("spec is nil")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(spec)
}

func readSpec(specFile string) ([]byte, error) {
	// Read the spec file
	data, err := os.ReadFile(specFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read spec file: %w", err)
	}
	return data, nil
}

func parseAndValidateSpec(data []byte) (*Spec, error) {
	var spec Spec

	// Unmarshal JSON into spec struct
	err := json.Unmarshal(data, &spec)
	if err != nil {
		return nil, fmt.Errorf("error parsing spec: %w", err)
	}
	spec.setDefaults()

	// Validate minimum structure
	if err := validateSpec(&spec); err != nil {
		return nil, err
	}

	return &spec, nil
}

func validateSpecFileValues(spec *Spec) error {
	for i, col := range spec.Columns {
		if col.Align != "" && !isValidAlignment(col.Align) {
			return fmt.Errorf("column %d: invalid align value: %q", i, col.Align)
		}

		for _, c := range col.Color.Default {
			if !isValidTextColor(c) {
				return fmt.Errorf("column %d: invalid color value: %q", i, c)
			}
		}
	}
	return nil
}

func readParseValidateSpec(specFile, envSpec, columns string) (*Spec, error) {
	if columns != "" {
		return convertColumnsToSpec(columns)
	}

	data := []byte(envSpec)
	if envSpec == "" {
		var err error
		data, err = readSpec(specFile)
		if err != nil {
			return nil, err
		}
	}
	spec, err := parseAndValidateSpec(data)
	if err != nil {
		return nil, err
	}

	if err := validateSpecFileValues(spec); err != nil {
		return nil, err
	}

	return spec, nil
}
