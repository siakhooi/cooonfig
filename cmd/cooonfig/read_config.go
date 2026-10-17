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

func ReadConfig(configFile string) (Config, error) {
	fmt.Printf("read config: %s\n", configFile)

	configRaw, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, fmt.Errorf("fail to read config file %s: %w", configFile, err)
	}
	cfg := Config{}
	if err := yaml.Unmarshal(configRaw, &cfg); err != nil {
		return Config{}, fmt.Errorf("fail to unmarshal config file %s: %w", configFile, err)
	}
	if len(cfg.Inputs) == 0 {
		return Config{}, fmt.Errorf("input files not config %s", configFile)
	}
	for _, input := range cfg.Inputs {
		if input.Type == "" {
			input.Type = "yaml"
		} else if input.Type != "yaml" {
			return Config{}, fmt.Errorf("invalid type %s: %s", input.Name, input.Type)
		}
	}
	return cfg, nil

}
