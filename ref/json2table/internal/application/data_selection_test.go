package application

import (
	"strings"
	"testing"
)

func TestSelectDataArraySuccess(t *testing.T) {
	fullData := map[string]interface{}{
		"items": []interface{}{
			map[string]interface{}{"name": "alice"},
			map[string]interface{}{"name": "bob"},
		},
	}

	got, err := selectDataArray("$.items", fullData)
	if err != nil {
		t.Fatalf("selectDataArray returned error: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("len(got) = %d, want 2", len(got))
	}
}

func TestSelectDataArrayInvalidJSONPath(t *testing.T) {
	fullData := map[string]interface{}{"items": []interface{}{1, 2}}

	_, err := selectDataArray("$.", fullData)
	if err == nil {
		t.Fatal("expected selectDataArray to return error")
	}
	if !strings.Contains(err.Error(), "error selecting data with jsonpath") {
		t.Fatalf("error = %q, want to contain %q", err.Error(), "error selecting data with jsonpath")
	}
}

func TestSelectDataArrayNotArray(t *testing.T) {
	fullData := map[string]interface{}{"item": map[string]interface{}{"name": "alice"}}

	_, err := selectDataArray("$.item", fullData)
	if err == nil {
		t.Fatal("expected selectDataArray to return error")
	}
	if !strings.Contains(err.Error(), "data selected with jsonpath is not an array") {
		t.Fatalf("error = %q, want to contain %q", err.Error(), "data selected with jsonpath is not an array")
	}
}
