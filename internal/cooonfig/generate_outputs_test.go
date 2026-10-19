package cooonfig

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeOutputTemplate(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestGenerateOutputs(t *testing.T) {
	t.Run("writes rendered files", func(t *testing.T) {
		templateDir := t.TempDir()
		outputDir := t.TempDir()
		writeOutputTemplate(t, templateDir, "out.txt", "hello {{ .name }}\n")

		err := generateOutputs(
			map[string]any{"name": "world"},
			[]output{{Template: "out.txt", OutFile: "nested/out.txt"}},
			templateDir,
			outputDir,
		)
		if err != nil {
			t.Fatalf("generateOutputs: %v", err)
		}

		path := filepath.Join(outputDir, "nested", "out.txt")
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "hello world\n" {
			t.Fatalf("content = %q", data)
		}
	})

	t.Run("empty outputs", func(t *testing.T) {
		if err := generateOutputs(map[string]any{}, nil, t.TempDir(), t.TempDir()); err != nil {
			t.Fatalf("generateOutputs: %v", err)
		}
	})

	t.Run("missing template", func(t *testing.T) {
		err := generateOutputs(
			map[string]any{},
			[]output{{Template: "missing.txt", OutFile: "out.txt"}},
			t.TempDir(),
			t.TempDir(),
		)
		if err == nil {
			t.Fatal("expected error")
		}
		if !strings.Contains(err.Error(), "get template") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("mkdir failure", func(t *testing.T) {
		templateDir := t.TempDir()
		outputDir := t.TempDir()
		writeOutputTemplate(t, templateDir, "out.txt", "ok\n")

		blocker := filepath.Join(outputDir, "blocker")
		if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}

		err := generateOutputs(
			map[string]any{},
			[]output{{Template: "out.txt", OutFile: "blocker/sub/out.txt"}},
			templateDir,
			outputDir,
		)
		if err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("template execute error", func(t *testing.T) {
		templateDir := t.TempDir()
		outputDir := t.TempDir()
		writeOutputTemplate(t, templateDir, "bad.txt", "{{ fail \"boom\" }}\n")

		err := generateOutputs(
			map[string]any{},
			[]output{{Template: "bad.txt", OutFile: "out.txt"}},
			templateDir,
			outputDir,
		)
		if err == nil {
			t.Fatal("expected error")
		}
		if !strings.Contains(err.Error(), "execute template") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("open file failure", func(t *testing.T) {
		templateDir := t.TempDir()
		outputDir := t.TempDir()
		writeOutputTemplate(t, templateDir, "out.txt", "ok\n")

		ro := filepath.Join(outputDir, "readonly")
		if err := os.Mkdir(ro, 0o555); err != nil {
			t.Fatal(err)
		}

		err := generateOutputs(
			map[string]any{},
			[]output{{Template: "out.txt", OutFile: "readonly/out.txt"}},
			templateDir,
			outputDir,
		)
		if err == nil {
			t.Fatal("expected error")
		}
		if !strings.Contains(err.Error(), "open file") {
			t.Fatalf("err = %v", err)
		}
	})
}
