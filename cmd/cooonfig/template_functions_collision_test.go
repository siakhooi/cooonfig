package main

import (
	"github.com/Masterminds/sprig/v3"
	"testing"
)

func TestSprigDoesNotDefineRequired(t *testing.T) {
	funcMap := sprig.TxtFuncMap()

	if _, exists := funcMap["required"]; exists {
		t.Fatal("Sprig now defined 'required', review cooonfig custom implementation if still needed")
	}
}
