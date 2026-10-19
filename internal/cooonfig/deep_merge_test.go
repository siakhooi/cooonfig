package cooonfig

import (
	"testing"
)

func TestDeepMerge(t *testing.T) {
	dest := map[string]any{
		"originString": "StringValue",
		"originInt":    3,
	}
	src := map[string]any{
		"newString": "StringValue",
		"newInt":    4,
	}
	deepMerge(dest, src)
	if dest["newString"] == nil {
		t.Fatalf("fail")
	}
}
