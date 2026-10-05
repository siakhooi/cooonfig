package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestRunHelp(t *testing.T) {
	for _, arg := range []string{"--help", "-h"} {
		t.Run(arg, func(t *testing.T) {
			output := captureStdout(t, func() {
				if err := run([]string{"cooonfig", arg}); err != nil {
					t.Fatalf("%s: %v", arg, err)
				}
			})
			if !strings.Contains(output, "--help") || !strings.Contains(output, "--version") {
				t.Fatalf("help output = %q", output)
			}
		})
	}
}

func TestRunVersion(t *testing.T) {
	for _, arg := range []string{"--version", "-v"} {
		t.Run(arg, func(t *testing.T) {
			output := captureStdout(t, func() {
				if err := run([]string{"cooonfig", arg}); err != nil {
					t.Fatalf("%s: %v", arg, err)
				}
			})
			if !strings.Contains(output, versionString()) {
				t.Fatalf("version output = %q", output)
			}
		})
	}
}

func TestRunNoArgsShowsHelp(t *testing.T) {
	output := captureStdout(t, func() {
		if err := run([]string{"cooonfig"}); err != nil {
			t.Fatalf("no args: %v", err)
		}
	})
	if !strings.Contains(output, "--help") || !strings.Contains(output, "--version") {
		t.Fatalf("expected help, got %q", output)
	}
}

func TestRunWithArgs(t *testing.T) {
	if err := run([]string{"cooonfig", "anything"}); err != nil {
		t.Fatalf("args: %v", err)
	}
}

func TestRunUnknownFlag(t *testing.T) {
	if err := run([]string{"cooonfig", "--spec"}); err == nil {
		t.Fatal("expected error for unknown flag")
	}
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	original := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	t.Cleanup(func() {
		os.Stdout = original
		_ = r.Close()
	})

	fn()
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatal(err)
	}
	return buf.String()
}
