package cooonfig

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type input struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
	Type string `yaml:"type"`
}
type output struct {
	Template string `yaml:"template"`
	OutFile  string `yaml:"outfile"`
}
type config struct {
	Inputs  []input  `yaml:"inputs"`
	Models  []string `yaml:"models"`
	Outputs []output `yaml:"outputs"`
}

const inputTypeYaml = "yaml"

func normalizeInputs(inputs []input) error {
	for i := range inputs {
		if inputs[i].Type == "" {
			inputs[i].Type = inputTypeYaml
		} else if inputs[i].Type != inputTypeYaml {
			return fmt.Errorf("invalid type %s: %s", inputs[i].Name, inputs[i].Type)
		}
	}
	return nil
}

func readConfig(configFile string) (config, error) {
	fmt.Printf("read config: %s\n", configFile)

	configRaw, err := os.ReadFile(configFile)
	if err != nil {
		return config{}, err
	}
	cfg := config{}
	if err := yaml.Unmarshal(configRaw, &cfg); err != nil {
		return config{}, fmt.Errorf("unmarshal %q: %w", configFile, err)
	}
	if len(cfg.Inputs) == 0 {
		return config{}, fmt.Errorf("missing input files %q", configFile)
	}
	if err := normalizeInputs(cfg.Inputs); err != nil {
		return config{}, err
	}

	return cfg, nil

}
