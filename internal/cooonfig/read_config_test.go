package cooonfig

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeConfigYAML(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "cooonfig.yaml")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func validConfigYAML() string {
	return `inputs:
  - name: app
    path: application.yaml
models:
  - models/core.yaml
outputs:
  - template: templates/out.txt
    outfile: out.txt
`
}

func TestReadConfig(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		path := writeConfigYAML(t, validConfigYAML())
		cfg, err := readConfig(path)
		if err != nil {
			t.Fatalf("readConfig: %v", err)
		}
		if len(cfg.Inputs) != 1 || cfg.Inputs[0].Name != "app" {
			t.Fatalf("inputs = %#v", cfg.Inputs)
		}
		if cfg.Inputs[0].Type != inputTypeYaml {
			t.Fatalf("type = %q, want yaml", cfg.Inputs[0].Type)
		}
		if len(cfg.Models) != 1 || len(cfg.Outputs) != 1 {
			t.Fatalf("cfg = %#v", cfg)
		}
	})

	t.Run("missing file", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "missing.yaml")
		_, err := readConfig(path)
		if err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("invalid yaml", func(t *testing.T) {
		path := writeConfigYAML(t, ":\n  bad\n  yaml\n")
		_, err := readConfig(path)
		if err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("missing inputs", func(t *testing.T) {
		path := writeConfigYAML(t, `models:
  - m.yaml
outputs:
  - template: t
    outfile: o
`)
		_, err := readConfig(path)
		if err == nil || !strings.Contains(err.Error(), "missing input files") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("missing models", func(t *testing.T) {
		path := writeConfigYAML(t, `inputs:
  - name: app
    path: app.yaml
outputs:
  - template: t
    outfile: o
`)
		_, err := readConfig(path)
		if err == nil || !strings.Contains(err.Error(), "missing model files") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("missing outputs", func(t *testing.T) {
		path := writeConfigYAML(t, `inputs:
  - name: app
    path: app.yaml
models:
  - m.yaml
`)
		_, err := readConfig(path)
		if err == nil || !strings.Contains(err.Error(), "missing output files") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("reserved input name", func(t *testing.T) {
		path := writeConfigYAML(t, `inputs:
  - name: model
    path: app.yaml
models:
  - m.yaml
outputs:
  - template: t
    outfile: o
`)
		_, err := readConfig(path)
		if err == nil || !strings.Contains(err.Error(), "reserved") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("duplicate input name", func(t *testing.T) {
		path := writeConfigYAML(t, `inputs:
  - name: app
    path: a.yaml
  - name: app
    path: b.yaml
models:
  - m.yaml
outputs:
  - template: t
    outfile: o
`)
		_, err := readConfig(path)
		if err == nil || !strings.Contains(err.Error(), "already exists") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("invalid input type", func(t *testing.T) {
		path := writeConfigYAML(t, `inputs:
  - name: app
    path: app.yaml
    type: json
models:
  - m.yaml
outputs:
  - template: t
    outfile: o
`)
		_, err := readConfig(path)
		if err == nil || !strings.Contains(err.Error(), "invalid type") {
			t.Fatalf("err = %v", err)
		}
	})
}

func TestIsReserved(t *testing.T) {
	if !isReserved(model) {
		t.Fatal("model should be reserved")
	}
	if isReserved("app") {
		t.Fatal("app should not be reserved")
	}
}

func TestNormalizeInputs(t *testing.T) {
	inputs := []input{{Name: "app", Path: "a.yaml", Type: ""}}
	if err := normalizeInputs(inputs); err != nil {
		t.Fatal(err)
	}
	if inputs[0].Type != inputTypeYaml {
		t.Fatalf("type = %q", inputs[0].Type)
	}
}
