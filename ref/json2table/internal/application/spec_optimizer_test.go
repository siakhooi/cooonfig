package application

import "testing"

func TestOptimizeSpecAppliesMinWidth(t *testing.T) {
	spec := &Spec{
		Columns: []Column{{Width: 3, MinWidth: 5, MaxWidth: 10}},
	}

	optimizeSpec(spec)

	if spec.Columns[0].Width != 5 {
		t.Fatalf("Width = %d, want %d", spec.Columns[0].Width, 5)
	}
}

func TestOptimizeSpecAppliesMaxWidth(t *testing.T) {
	spec := &Spec{
		Columns: []Column{{Width: 12, MinWidth: 1, MaxWidth: 8}},
	}

	optimizeSpec(spec)

	if spec.Columns[0].Width != 8 {
		t.Fatalf("Width = %d, want %d", spec.Columns[0].Width, 8)
	}
}

func TestOptimizeSpecKeepsWidthWhenWithinRange(t *testing.T) {
	spec := &Spec{
		Columns: []Column{{Width: 7, MinWidth: 5, MaxWidth: 10}},
	}

	optimizeSpec(spec)

	if spec.Columns[0].Width != 7 {
		t.Fatalf("Width = %d, want %d", spec.Columns[0].Width, 7)
	}
}

func TestOptimizeSpecIgnoresZeroBounds(t *testing.T) {
	spec := &Spec{
		Columns: []Column{{Width: 7, MinWidth: 0, MaxWidth: 0}},
	}

	optimizeSpec(spec)

	if spec.Columns[0].Width != 7 {
		t.Fatalf("Width = %d, want %d", spec.Columns[0].Width, 7)
	}
}

func TestOptimizeSpecProcessesAllColumns(t *testing.T) {
	spec := &Spec{
		Columns: []Column{
			{Title: "a", Width: 2, MinWidth: 4, MaxWidth: 8},
			{Title: "b", Width: 10, MinWidth: 2, MaxWidth: 6},
			{Title: "c", Width: 5, MinWidth: 1, MaxWidth: 9},
		},
	}

	optimizeSpec(spec)

	if spec.Columns[0].Width != 4 {
		t.Fatalf("Columns[0].Width = %d, want %d", spec.Columns[0].Width, 4)
	}
	if spec.Columns[1].Width != 6 {
		t.Fatalf("Columns[1].Width = %d, want %d", spec.Columns[1].Width, 6)
	}
	if spec.Columns[2].Width != 5 {
		t.Fatalf("Columns[2].Width = %d, want %d", spec.Columns[2].Width, 5)
	}
}
