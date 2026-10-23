package cooonfig

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func writeTemplateFile(t *testing.T, path, content string) {
	t.Helper()
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestGetTemplate(t *testing.T) {
	t.Run("parses and executes with sprig and required", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "nested", "tmpl.txt")
		writeTemplateFile(t, path, "{{ required \"need name\" .name | upper }}\n")

		tmpl, err := getTemplate(path)
		if err != nil {
			t.Fatalf("getTemplate: %v", err)
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, map[string]any{"name": "api"}); err != nil {
			t.Fatalf("execute: %v", err)
		}
		if buf.String() != "API\n" {
			t.Fatalf("output = %q", buf.String())
		}
	})

	t.Run("missing file", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "missing.txt")
		_, err := getTemplate(path)
		if err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("invalid template syntax", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "bad.txt")
		writeTemplateFile(t, path, "{{ .Name }\n")

		_, err := getTemplate(path)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func TestBuildTemplateFunctionMap(t *testing.T) {
	funcMap := buildTemplateFunctionMap()

	if _, ok := funcMap["required"]; !ok {
		t.Fatal("custom required missing")
	}
	if _, ok := funcMap["toYaml"]; !ok {
		t.Fatal("custom toYaml missing")
	}
	if _, ok := funcMap["upper"]; !ok {
		t.Fatal("sprig upper missing")
	}

	upper, ok := funcMap["upper"].(func(string) string)
	if !ok {
		t.Fatalf("upper type = %T", funcMap["upper"])
	}
	if upper("ab") != "AB" {
		t.Fatalf("upper = %q", upper("ab"))
	}
}
