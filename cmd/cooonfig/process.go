package main

import (
	"fmt"
)

func Process(configFile string) error {

	cfg, err := ReadConfig(configFile)
	if err != nil {
		return fmt.Errorf("fail to read Config %s: %w", configFile, err)
	}
	context, err := ReadInputs(cfg.Inputs)
	if err != nil {
		return fmt.Errorf("fail to read input files: %w", err)
	}

	model, err := RenderModel(context, cfg.Models)
	if err != nil {
		return fmt.Errorf("fail to render model: %w", err)
	}

	return GenerateOutputs(model, cfg.Outputs)

}
