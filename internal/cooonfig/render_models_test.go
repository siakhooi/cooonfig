package cooonfig

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func writeModelTemplate(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestRenderModel(t *testing.T) {
	t.Run("renders and merges models", func(t *testing.T) {
		dir := t.TempDir()
		writeModelTemplate(t, dir, "one.yaml", "x: 1\nnested:\n  a: 1\n")
		writeModelTemplate(t, dir, "two.yaml", "y: 2\nnested:\n  b: 2\n")

		context := map[string]any{
			"app": map[string]any{"name": "api"},
		}
		got, err := renderModel(context, []string{"one.yaml", "two.yaml"}, dir)
		if err != nil {
			t.Fatalf("renderModel: %v", err)
		}

		want := map[string]any{
			"x": 1,
			"y": 2,
			"nested": map[string]any{
				"a": 1,
				"b": 2,
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %#v, want %#v", got, want)
		}
	})

	t.Run("uses context in template", func(t *testing.T) {
		dir := t.TempDir()
		writeModelTemplate(t, dir, "ctx.yaml", "app_name: {{ .app.name }}\n")

		got, err := renderModel(map[string]any{
			"app": map[string]any{"name": "cooonfig"},
		}, []string{"ctx.yaml"}, dir)
		if err != nil {
			t.Fatalf("renderModel: %v", err)
		}
		if got["app_name"] != "cooonfig" {
			t.Fatalf("app_name = %v", got["app_name"])
		}
	})

	t.Run("missing model file", func(t *testing.T) {
		_, err := renderModel(map[string]any{}, []string{"missing.yaml"}, t.TempDir())
		if err == nil {
			t.Fatal("expected error")
		}
		if !strings.Contains(err.Error(), "get template") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("template execute error", func(t *testing.T) {
		dir := t.TempDir()
		writeModelTemplate(t, dir, "fail.yaml", "{{ fail \"boom\" }}\n")

		_, err := renderModel(map[string]any{}, []string{"fail.yaml"}, dir)
		if err == nil {
			t.Fatal("expected error")
		}
		if !strings.Contains(err.Error(), "execute template") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("invalid rendered yaml", func(t *testing.T) {
		dir := t.TempDir()
		writeModelTemplate(t, dir, "bad.yaml", "key: [unclosed\n")

		_, err := renderModel(map[string]any{}, []string{"bad.yaml"}, dir)
		if err == nil {
			t.Fatal("expected error")
		}
		if !strings.Contains(err.Error(), "unmarshal") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("empty model list", func(t *testing.T) {
		got, err := renderModel(map[string]any{}, nil, t.TempDir())
		if err != nil {
			t.Fatalf("renderModel: %v", err)
		}
		if len(got) != 0 {
			t.Fatalf("got %#v, want empty map", got)
		}
	})
}
