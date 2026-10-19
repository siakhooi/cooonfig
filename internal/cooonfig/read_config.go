package cooonfig

import (
	"fmt"
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

func isReserved(name string) bool {
	return name == model
}
func normalizeInputs(inputs []input) error {
	nameUsed := make(map[string]struct{}, len(inputs))
	for i := range inputs {
		if isReserved(inputs[i].Name) {
			return fmt.Errorf("input name is a reserved: %s", inputs[i].Name)
		}
		if _, exists := nameUsed[inputs[i].Name]; exists {
			return fmt.Errorf("input name already exists: %s", inputs[i].Name)
		}
		if inputs[i].Type == "" {
			inputs[i].Type = inputTypeYaml
		} else if inputs[i].Type != inputTypeYaml {
			return fmt.Errorf("invalid type %s: %s", inputs[i].Name, inputs[i].Type)
		}
		nameUsed[inputs[i].Name] = struct{}{}
	}
	return nil
}

func validateConfig(configFile string, cfg config) error {
	if len(cfg.Inputs) == 0 {
		return fmt.Errorf("missing input files %q", configFile)
	}
	if len(cfg.Models) == 0 {
		return fmt.Errorf("missing model files %q", configFile)
	}
	if len(cfg.Outputs) == 0 {
		return fmt.Errorf("missing output files %q", configFile)
	}
	if err := normalizeInputs(cfg.Inputs); err != nil {
		return err
	}

	return nil

}

func readConfig(configFile string) (config, error) {
	fmt.Printf("read config: %s\n", configFile)

	cfg := config{}
	if err := readYaml(configFile, &cfg); err != nil {
		return config{}, err
	}
	if err := validateConfig(configFile, cfg); err != nil {
		return config{}, err
	}

	return cfg, nil

}
