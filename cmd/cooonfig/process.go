package main

import (
	"fmt"
	"path/filepath"
)

func Process(configFile, inputDir, outputDir string) error {

	cfg, err := ReadConfig(configFile)
	if err != nil {
		return fmt.Errorf("read config %s: %w", configFile, err)
	}

	configDir := filepath.Dir(configFile)
	modelDir := configDir
	templateDir := configDir

	context, err := ReadInputs(cfg.Inputs, inputDir)
	if err != nil {
		return fmt.Errorf("read inputs: %w", err)
	}

	model, err := RenderModel(context, cfg.Models, modelDir)
	if err != nil {
		return fmt.Errorf("render model: %w", err)
	}

	return GenerateOutputs(model, cfg.Outputs, templateDir, outputDir)

}
