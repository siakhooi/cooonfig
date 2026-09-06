/*
Package application run the application
*/
package application

import (
	"fmt"

	"github.com/mattn/go-runewidth"
)

// Alignment represents text alignment in a column
type Alignment string

const (
	// AlignLeft aligns text to the left
	AlignLeft Alignment = "left"
	// AlignRight aligns text to the right
	AlignRight Alignment = "right"
	// AlignCenter aligns text to the center
	AlignCenter Alignment = "center"
)

// DefaultAlignment is the default text alignment (left)
var DefaultAlignment = AlignLeft

func formatAlignedTextParts(value string, width int, align Alignment) (string, string, string) {
	if width <= 0 {
		return "", "", ""
	}

	shortvalue := value
	if runewidth.StringWidth(value) > width {
		shortvalue = runewidth.Truncate(value, width, "")
	}
	prefix := ""
	suffix := ""
	visibleWidth := runewidth.StringWidth(shortvalue)
	if visibleWidth < width {
		padding := width - visibleWidth
		switch align {
		case AlignRight:
			prefix = fmt.Sprintf("%*s", padding, "")
		case AlignCenter:
			leftPad := padding / 2
			rightPad := padding - leftPad
			prefix = fmt.Sprintf("%*s", leftPad, "")
			suffix = fmt.Sprintf("%*s", rightPad, "")
		default: // AlignLeft
			suffix = fmt.Sprintf("%*s", padding, "")
		}
	}

	return prefix, shortvalue, suffix
}

func isValidAlignment(align Alignment) bool {
	validAlignments := map[Alignment]bool{
		AlignLeft:   true,
		AlignRight:  true,
		AlignCenter: true,
	}
	return validAlignments[align]
}
