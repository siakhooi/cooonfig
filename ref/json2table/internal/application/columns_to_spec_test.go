package application

import (
	"reflect"
	"testing"
)

func TestConvertColumnsToSpecBasic(t *testing.T) {
	spec, err := convertColumnsToSpec("id,name,age")
	if err != nil {
		t.Fatalf("convertColumnsToSpec returned error: %v", err)
	}
	if spec == nil {
		t.Fatal("convertColumnsToSpec returned nil spec")
	}
	if len(spec.Columns) != 3 {
		t.Fatalf("Columns length = %d, want 3", len(spec.Columns))
	}
	want := []string{"id", "name", "age"}
	for i, col := range spec.Columns {
		if !reflect.DeepEqual(col.Path, StringList{want[i]}) {
			t.Errorf("Column[%d].Path = %v, want %v", i, col.Path, StringList{want[i]})
		}
	}
}

func TestConvertColumnsToSpecTrimSpaces(t *testing.T) {
	spec, err := convertColumnsToSpec(" id , name , age ")
	if err != nil {
		t.Fatalf("convertColumnsToSpec returned error: %v", err)
	}
	want := []string{"id", "name", "age"}
	for i, col := range spec.Columns {
		if !reflect.DeepEqual(col.Path, StringList{want[i]}) {
			t.Errorf("Column[%d].Path = %v, want %v", i, col.Path, StringList{want[i]})
		}
	}
}

func TestSplitAndTrimCSV(t *testing.T) {
	input := " a , b ,c "
	want := []string{"a", "b", "c"}
	got := splitAndTrimCSV(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("splitAndTrimCSV(%q) = %v, want %v", input, got, want)
	}
}

func TestSplitCSV(t *testing.T) {
	input := "a,b,c"
	want := []string{"a", "b", "c"}
	got := splitCSV(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("splitCSV(%q) = %v, want %v", input, got, want)
	}
}

func TestTrimSpaces(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{"  hello  ", "hello"},
		{"\thello\t", "hello"},
		{"hello", "hello"},
		{"   ", ""},
	}
	for _, c := range cases {
		got := trimSpaces(c.in)
		if got != c.out {
			t.Errorf("trimSpaces(%q) = %q, want %q", c.in, got, c.out)
		}
	}
}
