package application

import "testing"

func TestUpdateColumnWidthSkipsInvalidAndNilPaths(t *testing.T) {
	spec := &Spec{
		Columns: []Column{
			{
				Path:  StringList{"missing", "name", "backup"},
				Width: 3,
			},
		},
	}

	item := map[string]interface{}{
		"name": "alice",
	}

	updateColumnWidth(spec, 0, spec.Columns[0], item)

	if got, want := spec.Columns[0].Width, 5; got != want {
		t.Fatalf("width = %d, want %d", got, want)
	}
}

func TestUpdateColumnWidthUsesFirstMatchingPathOnly(t *testing.T) {
	spec := &Spec{
		Columns: []Column{
			{
				Path:  StringList{"short", "long"},
				Width: 0,
			},
		},
	}

	item := map[string]interface{}{
		"short": "x",
		"long":  "this value is much longer",
	}

	updateColumnWidth(spec, 0, spec.Columns[0], item)

	if got, want := spec.Columns[0].Width, 1; got != want {
		t.Fatalf("width = %d, want %d", got, want)
	}
}

func TestUpdateColumnWidthContinuesAfterJSONPathError(t *testing.T) {
	spec := &Spec{
		Columns: []Column{
			{
				Path:  StringList{"$[", "name"},
				Width: 0,
			},
		},
	}

	item := map[string]interface{}{
		"name": "ok",
	}

	updateColumnWidth(spec, 0, spec.Columns[0], item)

	if got, want := spec.Columns[0].Width, len("ok"); got != want {
		t.Fatalf("width = %d, want %d", got, want)
	}
}

func TestAnalyseDataUpdatesAllColumnsToMaxWidth(t *testing.T) {
	spec := &Spec{
		Columns: []Column{
			{
				Path:  StringList{"name"},
				Width: 1,
			},
			{
				Path:  StringList{"title", "fallback"},
				Width: 2,
			},
		},
	}

	data := []interface{}{
		map[string]interface{}{"name": "Al", "title": "dev"},
		map[string]interface{}{"name": "Charlie", "fallback": "architect"},
		map[string]interface{}{"name": "Bob", "title": nil},
	}

	analyseData(spec, data)

	if got, want := spec.Columns[0].Width, len("Charlie"); got != want {
		t.Fatalf("name width = %d, want %d", got, want)
	}
	if got, want := spec.Columns[1].Width, len("architect"); got != want {
		t.Fatalf("title width = %d, want %d", got, want)
	}
}

func TestUpdateColumnWidthUsesDisplayWidthForAccents(t *testing.T) {
	spec := &Spec{
		Columns: []Column{
			{
				Path:  StringList{"name"},
				Width: 0,
			},
		},
	}

	item := map[string]interface{}{
		"name": "Padmé",
	}

	updateColumnWidth(spec, 0, spec.Columns[0], item)

	if got, want := spec.Columns[0].Width, 5; got != want {
		t.Fatalf("width = %d, want %d", got, want)
	}
}
