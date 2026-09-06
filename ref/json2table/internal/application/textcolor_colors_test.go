package application

import "testing"

func TestIsValidTextColorSupportedValues(t *testing.T) {
	supported := []string{
		ColorDefault,
		ColorRed,
		ColorGreen,
		ColorBlue,
		ColorYellow,
		ColorMagenta,
		ColorCyan,
		ColorWhite,
		ColorBlack,
		ColorHiRed,
		ColorHiGreen,
		ColorHiBlue,
		ColorHiYellow,
		ColorHiMagenta,
		ColorHiCyan,
		ColorHiWhite,
		ColorHiBlack,
		ColorBgRed,
		ColorBgGreen,
		ColorBgBlue,
		ColorBgYellow,
		ColorBgMagenta,
		ColorBgCyan,
		ColorBgWhite,
		ColorBgBlack,
		ColorHiBgRed,
		ColorHiBgGreen,
		ColorHiBgBlue,
		ColorHiBgYellow,
		ColorHiBgMagenta,
		ColorHiBgCyan,
		ColorHiBgWhite,
		ColorHiBgBlack,
		ColorBold,
		ColorFaint,
		ColorItalic,
		ColorUnderline,
		ColorBlinkSlow,
		ColorBlinkRapid,
		ColorReverseVideo,
		ColorConcealed,
		ColorCrossedOut,
	}

	for _, c := range supported {
		if !isValidTextColor(c) {
			t.Fatalf("isValidTextColor(%q) = false, want true", c)
		}
	}
}

func TestIsValidTextColorUnsupportedValues(t *testing.T) {
	unsupported := []string{"", "RED", "unknown", "bgpink", "bold "}

	for _, c := range unsupported {
		if isValidTextColor(c) {
			t.Fatalf("isValidTextColor(%q) = true, want false", c)
		}
	}
}

func TestSupportedColorMetaContainsKnownKeys(t *testing.T) {
	mustExist := []string{
		ColorDefault,
		ColorRed,
		ColorHiBlue,
		ColorBgYellow,
		ColorHiBgMagenta,
		ColorUnderline,
		ColorCrossedOut,
	}

	for _, key := range mustExist {
		if _, ok := supportedColorMeta[key]; !ok {
			t.Fatalf("supportedColorMeta is missing key %q", key)
		}
	}
}
