// Package cooonfig build a model from several configs, and render config/environment files
package cooonfig

import (
	"fmt"
	"path/filepath"
)

// Process builds a model from several configs, and render config/environment files
func Process(configFile, inputDir, outputDir string) error {

	cfg, err := readConfig(configFile)
	if err != nil {
		return fmt.Errorf("read config %s: %w", configFile, err)
	}

	configDir := filepath.Dir(configFile)
	modelDir := configDir
	templateDir := configDir

	context, err := readInputs(cfg.Inputs, inputDir)
	if err != nil {
		return fmt.Errorf("read inputs: %w", err)
	}

	model, err := renderModel(context, cfg.Models, modelDir)
	if err != nil {
		return fmt.Errorf("render model: %w", err)
	}

	return generateOutputs(model, cfg.Outputs, templateDir, outputDir)

}
