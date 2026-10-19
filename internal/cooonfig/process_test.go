package cooonfig

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func exampleRoot(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs(filepath.Join("..", "..", "example"))
	if err != nil {
		t.Fatal(err)
	}
	return root
}

func writeProcessFixture(t *testing.T, dir string, configYAML string, files map[string]string) string {
	t.Helper()
	for name, content := range files {
		path := filepath.Join(dir, name)
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	configPath := filepath.Join(dir, "cooonfig.yaml")
	if err := os.WriteFile(configPath, []byte(configYAML), 0o644); err != nil {
		t.Fatal(err)
	}
	return configPath
}

func TestProcess(t *testing.T) {
	t.Run("success with example", func(t *testing.T) {
		root := exampleRoot(t)
		outputDir := t.TempDir()

		err := Process(
			filepath.Join(root, "cooonfig.yaml"),
			root,
			outputDir,
		)
		if err != nil {
			t.Fatalf("Process: %v", err)
		}

		for _, path := range []string{
			"apache/config.json",
			"kafka/config.yaml",
		} {
			if _, err := os.Stat(filepath.Join(outputDir, path)); err != nil {
				t.Fatalf("expected output %s: %v", path, err)
			}
		}
	})

	t.Run("read config error", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "missing.yaml")
		err := Process(path, t.TempDir(), t.TempDir())
		if err == nil || !strings.Contains(err.Error(), "read config") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("read inputs error", func(t *testing.T) {
		dir := t.TempDir()
		configPath := writeProcessFixture(t, dir, `inputs:
  - name: app
    path: missing.yaml
models:
  - models/m.yaml
outputs:
  - template: templates/out.txt
    outfile: out.txt
`, map[string]string{
			"models/m.yaml":       "k: v\n",
			"templates/out.txt":   "ok\n",
		})

		err := Process(configPath, dir, t.TempDir())
		if err == nil || !strings.Contains(err.Error(), "read inputs") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("render model error", func(t *testing.T) {
		dir := t.TempDir()
		configPath := writeProcessFixture(t, dir, `inputs:
  - name: app
    path: app.yaml
models:
  - models/bad.yaml
outputs:
  - template: templates/out.txt
    outfile: out.txt
`, map[string]string{
			"app.yaml":          "name: api\n",
			"models/bad.yaml":   "{{ fail \"model boom\" }}\n",
			"templates/out.txt": "ok\n",
		})

		err := Process(configPath, dir, t.TempDir())
		if err == nil || !strings.Contains(err.Error(), "render model") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("generate outputs error", func(t *testing.T) {
		dir := t.TempDir()
		configPath := writeProcessFixture(t, dir, `inputs:
  - name: app
    path: app.yaml
models:
  - models/m.yaml
outputs:
  - template: templates/bad.txt
    outfile: out.txt
`, map[string]string{
			"app.yaml":          "name: api\n",
			"models/m.yaml":     "k: v\n",
			"templates/bad.txt": "{{ fail \"output boom\" }}\n",
		})

		err := Process(configPath, dir, t.TempDir())
		if err == nil || !strings.Contains(err.Error(), "execute template") {
			t.Fatalf("err = %v", err)
		}
	})
}
