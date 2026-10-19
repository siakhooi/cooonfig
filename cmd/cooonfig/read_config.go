package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Input struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
	Type string `yaml:"type"`
}
type Output struct {
	Template string `yaml:"template"`
	OutFile  string `yaml:"outfile"`
}
type Config struct {
	Inputs  []Input  `yaml:"inputs"`
	Models  []string `yaml:"models"`
	Outputs []Output `yaml:"outputs"`
}

const inputTypeYaml = "yaml"

func normalizeInputs(inputs []Input) error {
	for i := range inputs {
		if inputs[i].Type == "" {
			inputs[i].Type = inputTypeYaml
		} else if inputs[i].Type != inputTypeYaml {
			return fmt.Errorf("invalid type %s: %s", inputs[i].Name, inputs[i].Type)
		}
	}
	return nil
}

func ReadConfig(configFile string) (Config, error) {
	fmt.Printf("read config: %s\n", configFile)

	configRaw, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, err
	}
	cfg := Config{}
	if err := yaml.Unmarshal(configRaw, &cfg); err != nil {
		return Config{}, fmt.Errorf("unmarshal %q: %w", configFile, err)
	}
	if len(cfg.Inputs) == 0 {
		return Config{}, fmt.Errorf("missing input files %q", configFile)
	}
	if err := normalizeInputs(cfg.Inputs); err != nil {
		return Config{}, err
	}

	return cfg, nil

}
