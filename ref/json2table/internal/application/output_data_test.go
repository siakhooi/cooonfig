package application

import "testing"

func TestSelectFirstValueReturnsFirstMatchingNonNilValue(t *testing.T) {
	item := map[string]interface{}{
		"name": "alice",
		"age":  30,
	}

	got := selectFirstValue(StringList{"$.missing", "$.name", "$.age"}, item)
	if got != "alice" {
		t.Fatalf("selectFirstValue() = %v, want %v", got, "alice")
	}
}

func TestSelectFirstValueReturnsNilWhenNoPathMatches(t *testing.T) {
	item := map[string]interface{}{"name": "alice"}

	got := selectFirstValue(StringList{"$.missing", "$.other"}, item)
	if got != nil {
		t.Fatalf("selectFirstValue() = %v, want nil", got)
	}
}

func TestApplyURLPathReturnsOriginalWhenURLPathEmpty(t *testing.T) {
	printValue := "Alice"
	item := map[string]interface{}{"url": "https://example.com"}

	got := applyURLPath(printValue, "", item)
	if got != printValue {
		t.Fatalf("applyURLPath() = %q, want %q", got, printValue)
	}
}

func TestApplyURLPathReturnsOriginalWhenURLPathInvalid(t *testing.T) {
	printValue := "Alice"
	item := map[string]interface{}{"url": "https://example.com"}

	got := applyURLPath(printValue, "$.", item)
	if got != printValue {
		t.Fatalf("applyURLPath() = %q, want %q", got, printValue)
	}
}

func TestApplyURLPathAppliesLinkWhenURLExists(t *testing.T) {
	printValue := "Alice"
	item := map[string]interface{}{"url": "https://example.com"}

	got := applyURLPath(printValue, "$.url", item)
	want := applyLink(printValue, "https://example.com")
	if got != want {
		t.Fatalf("applyURLPath() = %q, want %q", got, want)
	}
}

func TestPrintDataPrintsRows(t *testing.T) {
	spec := &Spec{
		Columns: []Column{
			{Path: StringList{"$.name"}, Width: 5, Align: AlignLeft},
			{Path: StringList{"$.age"}, Width: 3, Align: AlignRight},
		},
	}
	dataArray := []interface{}{
		map[string]interface{}{"name": "bob", "age": 7},
		map[string]interface{}{"name": "alice", "age": 12},
	}

	got := captureStdout(t, func() {
		printData(dataArray, spec)
	})

	want := "bob     7 \nalice  12 \n"
	if got != want {
		t.Fatalf("printData output = %q, want %q", got, want)
	}
}
