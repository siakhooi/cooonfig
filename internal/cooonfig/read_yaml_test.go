package cooonfig

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadYaml(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "data.yaml")
		const content = "name: cooonfig\ncount: 2\n"
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}

		var got map[string]any
		if err := readYaml(path, &got); err != nil {
			t.Fatalf("readYaml: %v", err)
		}
		if got["name"] != "cooonfig" {
			t.Fatalf("name = %v", got["name"])
		}
		if got["count"] != 2 {
			t.Fatalf("count = %v", got["count"])
		}
	})

	t.Run("missing file", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "missing.yaml")
		var out map[string]any
		err := readYaml(path, &out)
		if err == nil {
			t.Fatal("expected error")
		}
		if !strings.Contains(err.Error(), "read file") || !strings.Contains(err.Error(), path) {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("invalid yaml", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "bad.yaml")
		if err := os.WriteFile(path, []byte(":\n  bad\n  indent"), 0o644); err != nil {
			t.Fatal(err)
		}

		var out map[string]any
		err := readYaml(path, &out)
		if err == nil {
			t.Fatal("expected error")
		}
		if !strings.Contains(err.Error(), "unmarshal") || !strings.Contains(err.Error(), path) {
			t.Fatalf("err = %v", err)
		}
	})
}

func TestUnmarshalYaml(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var got map[string]any
		if err := unmarshalYaml("inline.yaml", []byte("ok: true\n"), &got); err != nil {
			t.Fatal(err)
		}
		if got["ok"] != true {
			t.Fatalf("ok = %v", got["ok"])
		}
	})

	t.Run("invalid", func(t *testing.T) {
		var out map[string]any
		err := unmarshalYaml("broken.yaml", []byte("[unclosed"), &out)
		if err == nil {
			t.Fatal("expected error")
		}
		if !strings.Contains(err.Error(), "unmarshal") || !strings.Contains(err.Error(), "broken.yaml") {
			t.Fatalf("err = %v", err)
		}
	})
}
