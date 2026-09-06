/*
Package application run the application
*/
package application

import (
	"slices"

	"github.com/fatih/color"
)

func applyColor(originalValue, printValue string, textColor TextColorSpec) any {
	s := textColor.Default
	if textColor.Type == colorTypeConditional {
		for _, condition := range textColor.Conditions {
			if slices.Contains(condition.When, originalValue) {
				s = condition.Color
			}
		}
	}

	colors := make([]color.Attribute, 0, len(s))
	for _, c := range s {
		meta, ok := supportedColorMeta[c]
		if !ok {
			continue
		}
		colors = append(colors, meta.color)
	}
	if len(colors) == 0 {
		return printValue
	}
	return color.New(colors...).SprintFunc()(printValue)
}
