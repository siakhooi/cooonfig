package cooonfig

import (
	"fmt"
	"strings"
	"testing"
)

func TestTemplateFunctionRequired(t *testing.T) {
	const msg = "value is required"

	t.Run("nil", func(t *testing.T) {
		got, err := templateFunctionRequired(msg, nil)
		if got != nil {
			t.Fatalf("got %v, want nil", got)
		}
		if err == nil || err.Error() != msg {
			t.Fatalf("err = %v, want %q", err, msg)
		}
	})

	t.Run("empty string", func(t *testing.T) {
		got, err := templateFunctionRequired(msg, "")
		if got != nil {
			t.Fatalf("got %v, want nil", got)
		}
		if err == nil || err.Error() != msg {
			t.Fatalf("err = %v, want %q", err, msg)
		}
	})

	t.Run("non-empty string", func(t *testing.T) {
		const value = "ok"
		got, err := templateFunctionRequired(msg, value)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != value {
			t.Fatalf("got %v, want %q", got, value)
		}
	})

	t.Run("other types", func(t *testing.T) {
		for _, value := range []any{0, 42, false, true} {
			got, err := templateFunctionRequired(msg, value)
			if err != nil {
				t.Fatalf("value %v: %v", value, err)
			}
			if got != value {
				t.Fatalf("value %v: got %v", value, got)
			}
		}
	})
}

func TestTemplateFunctionToYaml(t *testing.T) {
	t.Run("map", func(t *testing.T) {
		got, err := templateFunctionToYaml(map[string]any{
			"name":  "api",
			"count": 2,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		const want = "count: 2\nname: api\n"
		if got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	})

	t.Run("slice", func(t *testing.T) {
		got, err := templateFunctionToYaml([]any{"a", "b"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		const want = "- a\n- b\n"
		if got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	})

	t.Run("nested", func(t *testing.T) {
		got, err := templateFunctionToYaml(map[string]any{
			"ports": []any{80, 443},
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		const want = "ports:\n    - 80\n    - 443\n"
		if got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	})

	t.Run("nil", func(t *testing.T) {
		got, err := templateFunctionToYaml(nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != "null\n" {
			t.Fatalf("got %q, want %q", got, "null\n")
		}
	})

	t.Run("marshal error", func(t *testing.T) {
		got, err := templateFunctionToYaml(failingYAML{})
		if err == nil || !strings.Contains(err.Error(), "cannot encode") {
			t.Fatalf("err = %v", err)
		}
		if got != "" {
			t.Fatalf("got %q, want empty string", got)
		}
	})
}

type failingYAML struct{}

func (failingYAML) MarshalYAML() (any, error) {
	return nil, fmt.Errorf("cannot encode")
}

func TestBuildTemplateFunctionMapToYaml(t *testing.T) {
	funcMap := buildTemplateFunctionMap()
	fn, ok := funcMap["toYaml"]
	if !ok {
		t.Fatal("toYaml not in template func map")
	}

	toYaml, ok := fn.(func(any) (string, error))
	if !ok {
		t.Fatalf("toYaml has unexpected type %T", fn)
	}

	got, err := toYaml("api")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "api\n" {
		t.Fatalf("got %q, want %q", got, "api\n")
	}
}

func TestBuildTemplateFunctionMapRequired(t *testing.T) {
	funcMap := buildTemplateFunctionMap()
	fn, ok := funcMap["required"]
	if !ok {
		t.Fatal("required not in template func map")
	}

	required, ok := fn.(func(string, any) (any, error))
	if !ok {
		t.Fatalf("required has unexpected type %T", fn)
	}

	got, err := required("missing", "")
	if err == nil || !strings.Contains(err.Error(), "missing") {
		t.Fatalf("err = %v", err)
	}
	if got != nil {
		t.Fatalf("got %v, want nil", got)
	}
}
