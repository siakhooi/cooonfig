package cooonfig

import (
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
