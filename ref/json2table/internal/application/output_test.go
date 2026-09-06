package application

import (
	"strings"
	"testing"
)

func TestPrintTableSuccess(t *testing.T) {
	spec := &Spec{
		DataPath: "$.items",
		Columns: []Column{
			{Path: StringList{"$.name"}, Title: "Name", Width: 4, Align: AlignLeft},
		},
	}
	fullData := map[string]interface{}{
		"items": []interface{}{
			map[string]interface{}{"name": "alice"},
			map[string]interface{}{"name": "bob"},
		},
	}

	got := captureStdout(t, func() {
		err := PrintTable(spec, fullData)
		if err != nil {
			t.Fatalf("PrintTable returned error: %v", err)
		}
	})

	want := "Name  \nalice \nbob   \n"
	if got != want {
		t.Fatalf("PrintTable output = %q, want %q", got, want)
	}
}

func TestPrintTableReturnsErrorWhenSelectedDataIsNotArray(t *testing.T) {
	spec := &Spec{DataPath: "$.item"}
	fullData := map[string]interface{}{"item": "not-array"}

	err := PrintTable(spec, fullData)
	if err == nil {
		t.Fatal("expected PrintTable to return error")
	}
	if !strings.Contains(err.Error(), "data selected with jsonpath is not an array") {
		t.Fatalf("error = %q, want to contain %q", err.Error(), "data selected with jsonpath is not an array")
	}
}
