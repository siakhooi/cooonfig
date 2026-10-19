package cooonfig

import (
	"reflect"
	"testing"
)

func TestDeepMerge(t *testing.T) {
	t.Run("adds missing keys", func(t *testing.T) {
		dst := map[string]any{
			"keep": "a",
		}
		src := map[string]any{
			"new": 1,
		}
		deepMerge(dst, src)
		want := map[string]any{
			"keep": "a",
			"new":  1,
		}
		if !reflect.DeepEqual(dst, want) {
			t.Fatalf("dst = %#v, want %#v", dst, want)
		}
	})

	t.Run("overrides scalar", func(t *testing.T) {
		dst := map[string]any{
			"x": "old",
		}
		src := map[string]any{
			"x": "new",
		}
		deepMerge(dst, src)
		if dst["x"] != "new" {
			t.Fatalf("x = %v, want new", dst["x"])
		}
	})

	t.Run("merges nested maps", func(t *testing.T) {
		dst := map[string]any{
			"nested": map[string]any{
				"a": 1,
				"b": 2,
			},
		}
		src := map[string]any{
			"nested": map[string]any{
				"b": 3,
				"c": 4,
			},
		}
		deepMerge(dst, src)
		wantNested := map[string]any{
			"a": 1,
			"b": 3,
			"c": 4,
		}
		nested, ok := dst["nested"].(map[string]any)
		if !ok {
			t.Fatalf("nested type = %T", dst["nested"])
		}
		if !reflect.DeepEqual(nested, wantNested) {
			t.Fatalf("nested = %#v, want %#v", nested, wantNested)
		}
	})

	t.Run("replaces map with scalar", func(t *testing.T) {
		dst := map[string]any{
			"k": map[string]any{"inner": true},
		}
		src := map[string]any{
			"k": "flat",
		}
		deepMerge(dst, src)
		if dst["k"] != "flat" {
			t.Fatalf("k = %v, want flat", dst["k"])
		}
	})

	t.Run("replaces scalar with map", func(t *testing.T) {
		dst := map[string]any{
			"k": "flat",
		}
		src := map[string]any{
			"k": map[string]any{"inner": true},
		}
		deepMerge(dst, src)
		want := map[string]any{"inner": true}
		got, ok := dst["k"].(map[string]any)
		if !ok {
			t.Fatalf("k type = %T", dst["k"])
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("k = %#v, want %#v", got, want)
		}
	})

	t.Run("empty source", func(t *testing.T) {
		dst := map[string]any{"a": 1}
		src := map[string]any{}
		deepMerge(dst, src)
		if !reflect.DeepEqual(dst, map[string]any{"a": 1}) {
			t.Fatalf("dst = %#v", dst)
		}
	})

	t.Run("deep nesting", func(t *testing.T) {
		dst := map[string]any{
			"l1": map[string]any{
				"l2": map[string]any{
					"leaf": "dst",
				},
			},
		}
		src := map[string]any{
			"l1": map[string]any{
				"l2": map[string]any{
					"leaf": "src",
					"extra": 9,
				},
			},
		}
		deepMerge(dst, src)
		l1 := dst["l1"].(map[string]any)
		l2 := l1["l2"].(map[string]any)
		if l2["leaf"] != "src" || l2["extra"] != 9 {
			t.Fatalf("l2 = %#v", l2)
		}
	})
}
