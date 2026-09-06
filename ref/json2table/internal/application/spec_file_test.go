package application

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateSpecNil(t *testing.T) {
	err := validateSpec(nil)
	if err == nil {
		t.Fatal("expected validateSpec to return error")
	}
	if !strings.Contains(err.Error(), "spec is nil") {
		t.Fatalf("error = %q, want to contain %q", err.Error(), "spec is nil")
	}
}

func TestValidateSpecValid(t *testing.T) {
	spec := &Spec{
		Columns: []Column{{Path: StringList{"$.name"}}},
	}

	err := validateSpec(spec)
	if err != nil {
		t.Fatalf("validateSpec returned error: %v", err)
	}
}

func TestReadSpecSuccess(t *testing.T) {
	tmpDir := t.TempDir()
	specFile := filepath.Join(tmpDir, "spec.json")
	want := []byte(`{"columns":[{"path":"$.name"}]}`)
	if err := os.WriteFile(specFile, want, 0o644); err != nil {
		t.Fatalf("failed to create spec file: %v", err)
	}

	got, err := readSpec(specFile)
	if err != nil {
		t.Fatalf("readSpec returned error: %v", err)
	}
	if string(got) != string(want) {
		t.Fatalf("readSpec returned %q, want %q", string(got), string(want))
	}
}

func TestReadSpecMissingFile(t *testing.T) {
	_, err := readSpec(filepath.Join(t.TempDir(), "missing.json"))
	if err == nil {
		t.Fatal("expected readSpec to return error")
	}
	if !strings.Contains(err.Error(), "cannot read spec file") {
		t.Fatalf("error = %q, want to contain %q", err.Error(), "cannot read spec file")
	}
}

func TestParseAndValidateSpecValid(t *testing.T) {
	spec, err := parseAndValidateSpec([]byte(`{"columns":[{"path":"$.person.name"}]}`))
	if err != nil {
		t.Fatalf("parseAndValidateSpec returned error: %v", err)
	}
	if spec == nil {
		t.Fatal("parseAndValidateSpec returned nil spec")
	}
	if spec.DataPath != "$" {
		t.Fatalf("DataPath = %q, want %q", spec.DataPath, "$")
	}
	if len(spec.Columns) != 1 {
		t.Fatalf("Columns length = %d, want 1", len(spec.Columns))
	}
	if spec.Columns[0].Title != "name" {
		t.Fatalf("Columns[0].Title = %q, want %q", spec.Columns[0].Title, "name")
	}
}

func TestParseAndValidateSpecInvalidJSON(t *testing.T) {
	_, err := parseAndValidateSpec([]byte(`{"columns":`))
	if err == nil {
		t.Fatal("expected parseAndValidateSpec to return error")
	}
	if !strings.Contains(err.Error(), "error parsing spec") {
		t.Fatalf("error = %q, want to contain %q", err.Error(), "error parsing spec")
	}
}

func TestParseAndValidateSpecValidationError(t *testing.T) {
	_, err := parseAndValidateSpec([]byte(`{"columns":[]}`))
	if err == nil {
		t.Fatal("expected parseAndValidateSpec to return validation error")
	}
}

func TestValidateSpecFileValuesInvalidAlign(t *testing.T) {
	spec := &Spec{
		Columns: []Column{{
			Path:  StringList{"$.name"},
			Align: Alignment("middle"),
			Color: TextColorSpec{Type: colorTypeFixed, Default: StringList{ColorRed}},
		}},
	}

	err := validateSpecFileValues(spec)
	if err == nil {
		t.Fatal("expected validateSpecFileValues to return error")
	}
	if !strings.Contains(err.Error(), "invalid align value") {
		t.Fatalf("error = %q, want to contain %q", err.Error(), "invalid align value")
	}
}

func TestValidateSpecFileValuesInvalidColor(t *testing.T) {
	spec := &Spec{
		Columns: []Column{{
			Path:  StringList{"$.name"},
			Align: AlignLeft,
			Color: TextColorSpec{Type: colorTypeFixed, Default: StringList{"unknown-color"}},
		}},
	}

	err := validateSpecFileValues(spec)
	if err == nil {
		t.Fatal("expected validateSpecFileValues to return error")
	}
	if !strings.Contains(err.Error(), "invalid color value") {
		t.Fatalf("error = %q, want to contain %q", err.Error(), "invalid color value")
	}
}

func TestValidateSpecFileValuesValid(t *testing.T) {
	spec := &Spec{
		Columns: []Column{{
			Path:  StringList{"$.name"},
			Align: AlignCenter,
			Color: TextColorSpec{Type: colorTypeFixed, Default: StringList{ColorBlue, ColorBold}},
		}},
	}

	err := validateSpecFileValues(spec)
	if err != nil {
		t.Fatalf("validateSpecFileValues returned error: %v", err)
	}
}

func TestReadParseValidateSpecUsesEnvSpecWhenProvided(t *testing.T) {
	envSpec := `{"columns":[{"path":"$.name","color":"green"}]}`
	spec, err := readParseValidateSpec("/path/does/not/exist/spec.json", envSpec, "")
	if err != nil {
		t.Fatalf("ReadParseValidateSpec returned error: %v", err)
	}
	if spec == nil {
		t.Fatal("ReadParseValidateSpec returned nil spec")
	}
	if len(spec.Columns) != 1 {
		t.Fatalf("Columns length = %d, want 1", len(spec.Columns))
	}
}

func TestReadParseValidateSpecReadsFromFile(t *testing.T) {
	tmpDir := t.TempDir()
	specFile := filepath.Join(tmpDir, "spec.json")
	content := []byte(`{"dataPath":"$.items","columns":[{"path":"$.name","align":"right","color":"red"}]}`)
	if err := os.WriteFile(specFile, content, 0o644); err != nil {
		t.Fatalf("failed to write spec file: %v", err)
	}

	spec, err := readParseValidateSpec(specFile, "", "")
	if err != nil {
		t.Fatalf("ReadParseValidateSpec returned error: %v", err)
	}
	if spec.DataPath != "$.items" {
		t.Fatalf("DataPath = %q, want %q", spec.DataPath, "$.items")
	}
	if spec.Columns[0].Align != AlignRight {
		t.Fatalf("Columns[0].Align = %q, want %q", spec.Columns[0].Align, AlignRight)
	}
}

func TestReadParseValidateSpecReadFileError(t *testing.T) {
	_, err := readParseValidateSpec(filepath.Join(t.TempDir(), "missing.json"), "", "")
	if err == nil {
		t.Fatal("expected ReadParseValidateSpec to return error")
	}
	if !strings.Contains(err.Error(), "cannot read spec file") {
		t.Fatalf("error = %q, want to contain %q", err.Error(), "cannot read spec file")
	}
}

func TestReadParseValidateSpecInvalidColorValue(t *testing.T) {
	envSpec := `{"columns":[{"path":"$.name","color":"unknown-color"}]}`
	_, err := readParseValidateSpec("", envSpec, "")
	if err == nil {
		t.Fatal("expected ReadParseValidateSpec to return error")
	}
	if !strings.Contains(err.Error(), "invalid color value") {
		t.Fatalf("error = %q, want to contain %q", err.Error(), "invalid color value")
	}
}

func TestReadParseValidateSpecColumnsArg(t *testing.T) {
	spec, err := readParseValidateSpec("", "", "id,name,age")
	if err != nil {
		t.Fatalf("ReadParseValidateSpec with columns arg returned error: %v", err)
	}
	if spec == nil {
		t.Fatal("ReadParseValidateSpec with columns arg returned nil spec")
	}
	if len(spec.Columns) != 3 {
		t.Fatalf("Columns length = %d, want 3", len(spec.Columns))
	}
	want := []string{"id", "name", "age"}
	for i, col := range spec.Columns {
		if len(col.Path) != 1 || col.Path[0] != want[i] {
			t.Errorf("Column[%d].Path = %v, want %q", i, col.Path, want[i])
		}
	}
}
