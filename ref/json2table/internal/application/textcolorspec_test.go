package application

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestTextColorSpecUnmarshalString(t *testing.T) {
	var spec TextColorSpec
	err := json.Unmarshal([]byte(`"red"`), &spec)
	if err != nil {
		t.Fatalf("json.Unmarshal returned error: %v", err)
	}

	if spec.Type != colorTypeFixed {
		t.Fatalf("Type = %v, want %v", spec.Type, colorTypeFixed)
	}
	if !reflect.DeepEqual(spec.Default, StringList{"red"}) {
		t.Fatalf("Default = %v, want %v", spec.Default, StringList{"red"})
	}
	if len(spec.Conditions) != 0 {
		t.Fatalf("Conditions length = %d, want 0", len(spec.Conditions))
	}
}

func TestTextColorSpecUnmarshalStringArray(t *testing.T) {
	var spec TextColorSpec
	err := json.Unmarshal([]byte(`["red", "bold"]`), &spec)
	if err != nil {
		t.Fatalf("json.Unmarshal returned error: %v", err)
	}

	if spec.Type != colorTypeFixed {
		t.Fatalf("Type = %v, want %v", spec.Type, colorTypeFixed)
	}
	if !reflect.DeepEqual(spec.Default, StringList{"red", "bold"}) {
		t.Fatalf("Default = %v, want %v", spec.Default, StringList{"red", "bold"})
	}
	if len(spec.Conditions) != 0 {
		t.Fatalf("Conditions length = %d, want 0", len(spec.Conditions))
	}
}

func TestTextColorSpecUnmarshalObjectWithDefault(t *testing.T) {
	input := `{
		"default": ["yellow"],
		"conditions": [
			{"when": ["< 0"], "color": ["red"]},
			{"when": [">= 0"], "color": ["green"]}
		]
	}`

	var spec TextColorSpec
	err := json.Unmarshal([]byte(input), &spec)
	if err != nil {
		t.Fatalf("json.Unmarshal returned error: %v", err)
	}

	if spec.Type != colorTypeConditional {
		t.Fatalf("Type = %v, want %v", spec.Type, colorTypeConditional)
	}
	if !reflect.DeepEqual(spec.Default, StringList{"yellow"}) {
		t.Fatalf("Default = %v, want %v", spec.Default, StringList{"yellow"})
	}
	if len(spec.Conditions) != 2 {
		t.Fatalf("Conditions length = %d, want 2", len(spec.Conditions))
	}
	if !reflect.DeepEqual(spec.Conditions[0].When, StringList{"< 0"}) {
		t.Fatalf("Conditions[0].When = %v, want %v", spec.Conditions[0].When, StringList{"< 0"})
	}
	if !reflect.DeepEqual(spec.Conditions[0].Color, StringList{"red"}) {
		t.Fatalf("Conditions[0].Color = %v, want %v", spec.Conditions[0].Color, StringList{"red"})
	}
	if !reflect.DeepEqual(spec.Conditions[1].When, StringList{">= 0"}) {
		t.Fatalf("Conditions[1].When = %v, want %v", spec.Conditions[1].When, StringList{">= 0"})
	}
	if !reflect.DeepEqual(spec.Conditions[1].Color, StringList{"green"}) {
		t.Fatalf("Conditions[1].Color = %v, want %v", spec.Conditions[1].Color, StringList{"green"})
	}
}

func TestTextColorSpecUnmarshalObjectWithoutDefaultUsesColorDefault(t *testing.T) {
	input := `{
		"conditions": [
			{"when": ["== null"], "color": ["blue"]}
		]
	}`

	var spec TextColorSpec
	err := json.Unmarshal([]byte(input), &spec)
	if err != nil {
		t.Fatalf("json.Unmarshal returned error: %v", err)
	}

	if spec.Type != colorTypeConditional {
		t.Fatalf("Type = %v, want %v", spec.Type, colorTypeConditional)
	}
	if !reflect.DeepEqual(spec.Default, StringList{ColorDefault}) {
		t.Fatalf("Default = %v, want %v", spec.Default, StringList{ColorDefault})
	}
	if len(spec.Conditions) != 1 {
		t.Fatalf("Conditions length = %d, want 1", len(spec.Conditions))
	}
	if !reflect.DeepEqual(spec.Conditions[0].Color, StringList{"blue"}) {
		t.Fatalf("Conditions[0].Color = %v, want %v", spec.Conditions[0].Color, StringList{"blue"})
	}
}

func TestTextColorSpecUnmarshalInvalidSpec(t *testing.T) {
	var spec TextColorSpec
	err := json.Unmarshal([]byte(`123`), &spec)
	if err == nil {
		t.Fatal("expected json.Unmarshal to return error")
	}
	if !strings.Contains(err.Error(), "invalid color specification") {
		t.Fatalf("error = %q, want to contain %q", err.Error(), "invalid color specification")
	}
}
