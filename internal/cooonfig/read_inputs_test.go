package cooonfig

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func writeInputYAML(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return name
}

func TestReadInputs(t *testing.T) {
	t.Run("loads named inputs", func(t *testing.T) {
		dir := t.TempDir()
		writeInputYAML(t, dir, "app.yaml", "service: api\n")
		writeInputYAML(t, dir, "env.yaml", "region: west\n")

		got, err := readInputs([]input{
			{Name: "app", Path: "app.yaml"},
			{Name: "env", Path: "env.yaml"},
		}, dir)
		if err != nil {
			t.Fatalf("readInputs: %v", err)
		}

		want := map[string]any{
			"app": map[string]any{"service": "api"},
			"env": map[string]any{"region": "west"},
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %#v, want %#v", got, want)
		}
	})

	t.Run("empty list", func(t *testing.T) {
		got, err := readInputs(nil, t.TempDir())
		if err != nil {
			t.Fatalf("readInputs: %v", err)
		}
		if len(got) != 0 {
			t.Fatalf("got %#v, want empty map", got)
		}
	})

	t.Run("missing file", func(t *testing.T) {
		dir := t.TempDir()
		_, err := readInputs([]input{
			{Name: "app", Path: "missing.yaml"},
		}, dir)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}
