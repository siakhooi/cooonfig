package cooonfig

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func unmarshalYaml(yamlPath string, data []byte, out any) error {
	if err := yaml.Unmarshal(data, out); err != nil {
		return fmt.Errorf("unmarshal %q: %w", yamlPath, err)
	}
	return nil

}
func readYaml(yamlPath string, out any) error {
	data, err := os.ReadFile(yamlPath)
	if err != nil {
		return fmt.Errorf("read file %q: %w", yamlPath, err)
	}
	return unmarshalYaml(yamlPath, data, out)
}
