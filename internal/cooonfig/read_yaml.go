package cooonfig

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func readYaml(yamlPath string, out any) error {
	data, err := os.ReadFile(yamlPath)
	if err != nil {
		return fmt.Errorf("read file %q: %w", yamlPath, err)
	}
	if err := yaml.Unmarshal(data, out); err != nil {
		return fmt.Errorf("unmarshal %q: %w", yamlPath, err)
	}
	return nil
}
