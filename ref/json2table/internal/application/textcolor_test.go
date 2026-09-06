package application

import (
	"testing"

	"github.com/fatih/color"
)

func withColorEnabled(t *testing.T) {
	t.Helper()
	originalNoColor := color.NoColor
	color.NoColor = false
	t.Cleanup(func() {
		color.NoColor = originalNoColor
	})
}

func TestGetColoredFixedValidColor(t *testing.T) {
	withColorEnabled(t)

	printValue := "hello"
	spec := TextColorSpec{Type: colorTypeFixed, Default: StringList{ColorRed}}

	got := applyColor("ignored", printValue, spec)
	expected := color.New(color.FgRed).SprintFunc()(printValue)
	if got != expected {
		t.Fatalf("GetColored() = %q, want %q", got, expected)
	}
}

func TestGetColoredUnsupportedColorReturnsPlainValue(t *testing.T) {
	printValue := "hello"
	spec := TextColorSpec{Type: colorTypeFixed, Default: StringList{"unknown-color"}}

	got := applyColor("ignored", printValue, spec)
	if got != printValue {
		t.Fatalf("GetColored() = %q, want %q", got, printValue)
	}
}

func TestGetColoredConditionalMatchUsesConditionColor(t *testing.T) {
	withColorEnabled(t)

	printValue := "hello"
	spec := TextColorSpec{
		Type:    colorTypeConditional,
		Default: StringList{ColorGreen},
		Conditions: []struct {
			When  StringList `json:"when"`
			Color StringList `json:"color"`
		}{
			{When: StringList{"match"}, Color: StringList{ColorRed}},
		},
	}

	got := applyColor("match", printValue, spec)
	expected := color.New(color.FgRed).SprintFunc()(printValue)
	if got != expected {
		t.Fatalf("GetColored() = %q, want %q", got, expected)
	}
}

func TestGetColoredConditionalNoMatchUsesDefaultColor(t *testing.T) {
	withColorEnabled(t)

	printValue := "hello"
	spec := TextColorSpec{
		Type:    colorTypeConditional,
		Default: StringList{ColorGreen},
		Conditions: []struct {
			When  StringList `json:"when"`
			Color StringList `json:"color"`
		}{
			{When: StringList{"match"}, Color: StringList{ColorRed}},
		},
	}

	got := applyColor("no-match", printValue, spec)
	expected := color.New(color.FgGreen).SprintFunc()(printValue)
	if got != expected {
		t.Fatalf("GetColored() = %q, want %q", got, expected)
	}
}

func TestGetColoredConditionalLastMatchWins(t *testing.T) {
	withColorEnabled(t)

	printValue := "hello"
	spec := TextColorSpec{
		Type:    colorTypeConditional,
		Default: StringList{ColorGreen},
		Conditions: []struct {
			When  StringList `json:"when"`
			Color StringList `json:"color"`
		}{
			{When: StringList{"match"}, Color: StringList{ColorRed}},
			{When: StringList{"match"}, Color: StringList{ColorBlue}},
		},
	}

	got := applyColor("match", printValue, spec)
	expected := color.New(color.FgBlue).SprintFunc()(printValue)
	if got != expected {
		t.Fatalf("GetColored() = %q, want %q", got, expected)
	}
}
