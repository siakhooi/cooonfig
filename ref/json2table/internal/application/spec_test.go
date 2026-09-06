package application

import (
	"math"
	"reflect"
	"testing"
)

func TestColumnSetDefaultsDerivesTitleFromPath(t *testing.T) {
	column := Column{Path: StringList{"$.user.name"}}

	column.setDefaults()

	if column.Title != "name" {
		t.Fatalf("Title = %q, want %q", column.Title, "name")
	}
	if column.Width != len("name") {
		t.Fatalf("Width = %d, want %d", column.Width, len("name"))
	}
}

func TestColumnSetDefaultsUsesPathWhenNoDot(t *testing.T) {
	column := Column{Path: StringList{"name"}}

	column.setDefaults()

	if column.Title != "name" {
		t.Fatalf("Title = %q, want %q", column.Title, "name")
	}
}

func TestColumnSetDefaultsPreservesProvidedValues(t *testing.T) {
	customColor := TextColorSpec{Type: colorTypeFixed, Default: StringList{ColorRed}}
	column := Column{
		Path:     StringList{"$.name"},
		Title:    "User",
		MaxWidth: 12,
		Align:    AlignRight,
		Color:    customColor,
	}

	column.setDefaults()

	if column.Title != "User" {
		t.Fatalf("Title = %q, want %q", column.Title, "User")
	}
	if column.Width != len("User") {
		t.Fatalf("Width = %d, want %d", column.Width, len("User"))
	}
	if column.MaxWidth != 12 {
		t.Fatalf("MaxWidth = %d, want %d", column.MaxWidth, 12)
	}
	if column.Align != AlignRight {
		t.Fatalf("Align = %q, want %q", column.Align, AlignRight)
	}
	if !reflect.DeepEqual(column.Color, customColor) {
		t.Fatalf("Color = %#v, want %#v", column.Color, customColor)
	}
}

func TestColumnSetDefaultsAssignsFallbacks(t *testing.T) {
	column := Column{Path: StringList{"$.name"}}

	column.setDefaults()

	if column.MaxWidth != math.MaxInt {
		t.Fatalf("MaxWidth = %d, want %d", column.MaxWidth, math.MaxInt)
	}
	if column.Align != DefaultAlignment {
		t.Fatalf("Align = %q, want %q", column.Align, DefaultAlignment)
	}
	if !reflect.DeepEqual(column.Color, DefaultTextColor) {
		t.Fatalf("Color = %#v, want %#v", column.Color, DefaultTextColor)
	}
}

func TestSpecSetDefaultsSetsDataPathAndColumns(t *testing.T) {
	spec := Spec{
		Columns: []Column{
			{Path: StringList{"$.name"}},
			{Path: StringList{"$.age"}, Title: "Age", Align: AlignCenter, Color: TextColorSpec{Type: colorTypeFixed, Default: StringList{ColorBlue}}},
		},
	}

	spec.setDefaults()

	if spec.DataPath != "$" {
		t.Fatalf("DataPath = %q, want %q", spec.DataPath, "$")
	}
	if spec.Columns[0].Title != "name" {
		t.Fatalf("Columns[0].Title = %q, want %q", spec.Columns[0].Title, "name")
	}
	if spec.Columns[0].Align != DefaultAlignment {
		t.Fatalf("Columns[0].Align = %q, want %q", spec.Columns[0].Align, DefaultAlignment)
	}
	if spec.Columns[1].Title != "Age" {
		t.Fatalf("Columns[1].Title = %q, want %q", spec.Columns[1].Title, "Age")
	}
	if spec.Columns[1].Align != AlignCenter {
		t.Fatalf("Columns[1].Align = %q, want %q", spec.Columns[1].Align, AlignCenter)
	}
}

func TestSpecSetDefaultsPreservesDataPath(t *testing.T) {
	spec := Spec{
		DataPath: "$.items",
		Columns:  []Column{{Path: StringList{"$.name"}}},
	}

	spec.setDefaults()

	if spec.DataPath != "$.items" {
		t.Fatalf("DataPath = %q, want %q", spec.DataPath, "$.items")
	}
}
