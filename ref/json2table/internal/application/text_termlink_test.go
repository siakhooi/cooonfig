package application

import (
	"testing"

	"github.com/savioxavier/termlink"
)

func TestGetLink_EmptyURLReturnsText(t *testing.T) {
	text := "hello"
	got := applyLink(text, "")
	if got != text {
		t.Fatalf("GetLink(%q, %q) = %q, want %q", text, "", got, text)
	}
}

func TestGetLink_WithURLReturnsTerminalLink(t *testing.T) {
	text := "repo"
	url := "https://example.com"
	want := termlink.Link(text, url)

	got := applyLink(text, url)
	if got != want {
		t.Fatalf("GetLink(%q, %q) = %q, want %q", text, url, got, want)
	}
}
