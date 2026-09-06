package application

import "testing"

func TestFormatAlignedTextParts(t *testing.T) {
	tests := []struct {
		name       string
		value      string
		width      int
		align      Alignment
		wantPrefix string
		wantValue  string
		wantSuffix string
	}{
		{
			name:       "left align pads suffix",
			value:      "ab",
			width:      5,
			align:      AlignLeft,
			wantPrefix: "",
			wantValue:  "ab",
			wantSuffix: "   ",
		},
		{
			name:       "right align pads prefix",
			value:      "ab",
			width:      5,
			align:      AlignRight,
			wantPrefix: "   ",
			wantValue:  "ab",
			wantSuffix: "",
		},
		{
			name:       "center align splits padding",
			value:      "ab",
			width:      5,
			align:      AlignCenter,
			wantPrefix: " ",
			wantValue:  "ab",
			wantSuffix: "  ",
		},
		{
			name:       "center align odd padding with longer text",
			value:      "cat",
			width:      8,
			align:      AlignCenter,
			wantPrefix: "  ",
			wantValue:  "cat",
			wantSuffix: "   ",
		},
		{
			name:       "truncates when value exceeds width",
			value:      "abcdef",
			width:      4,
			align:      AlignLeft,
			wantPrefix: "",
			wantValue:  "abcd",
			wantSuffix: "",
		},
		{
			name:       "exact width has no padding",
			value:      "abcd",
			width:      4,
			align:      AlignRight,
			wantPrefix: "",
			wantValue:  "abcd",
			wantSuffix: "",
		},
		{
			name:       "accented characters use display width for padding",
			value:      "Padmé",
			width:      8,
			align:      AlignLeft,
			wantPrefix: "",
			wantValue:  "Padmé",
			wantSuffix: "   ",
		},
		{
			name:       "unicode truncation does not split multi-byte rune",
			value:      "Ric Olié",
			width:      7,
			align:      AlignLeft,
			wantPrefix: "",
			wantValue:  "Ric Oli",
			wantSuffix: "",
		},
		{
			name:       "zero width yields empty value",
			value:      "abcd",
			width:      0,
			align:      AlignLeft,
			wantPrefix: "",
			wantValue:  "",
			wantSuffix: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotPrefix, gotValue, gotSuffix := formatAlignedTextParts(tc.value, tc.width, tc.align)

			if gotPrefix != tc.wantPrefix || gotValue != tc.wantValue || gotSuffix != tc.wantSuffix {
				t.Fatalf(
					"FormatAlignedTextParts(%q, %d, %q) = (%q, %q, %q), want (%q, %q, %q)",
					tc.value,
					tc.width,
					tc.align,
					gotPrefix,
					gotValue,
					gotSuffix,
					tc.wantPrefix,
					tc.wantValue,
					tc.wantSuffix,
				)
			}
		})
	}
}

func TestIsValidAlignment(t *testing.T) {
	tests := []struct {
		name  string
		align Alignment
		want  bool
	}{
		{
			name:  "left alignment is valid",
			align: AlignLeft,
			want:  true,
		},
		{
			name:  "right alignment is valid",
			align: AlignRight,
			want:  true,
		},
		{
			name:  "center alignment is valid",
			align: AlignCenter,
			want:  true,
		},
		{
			name:  "empty alignment is invalid",
			align: "",
			want:  false,
		},
		{
			name:  "unknown alignment is invalid",
			align: Alignment("justify"),
			want:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isValidAlignment(tc.align)
			if got != tc.want {
				t.Fatalf("isValidAlignment(%q) = %v, want %v", tc.align, got, tc.want)
			}
		})
	}
}
