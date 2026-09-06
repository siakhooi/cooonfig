package application

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	originalStdout := os.Stdout
	readPipe, writePipe, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdout pipe: %v", err)
	}
	os.Stdout = writePipe
	t.Cleanup(func() {
		os.Stdout = originalStdout
	})

	fn()

	if err := writePipe.Close(); err != nil {
		t.Fatalf("failed to close write pipe: %v", err)
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, readPipe); err != nil {
		t.Fatalf("failed to read captured stdout: %v", err)
	}
	if err := readPipe.Close(); err != nil {
		t.Fatalf("failed to close read pipe: %v", err)
	}

	return buf.String()
}

func TestPrintHeaderPrintsAlignedColumns(t *testing.T) {
	columns := []Column{
		{Title: "ID", Width: 4, Align: AlignRight},
		{Title: "Name", Width: 6, Align: AlignLeft},
	}

	got := captureStdout(t, func() {
		printHeader(columns)
	})

	want := "  ID Name   \n"
	if got != want {
		t.Fatalf("printHeader output = %q, want %q", got, want)
	}
}

func TestPrintHeaderWithNoColumnsPrintsNewlineOnly(t *testing.T) {
	got := captureStdout(t, func() {
		printHeader(nil)
	})

	want := "\n"
	if got != want {
		t.Fatalf("printHeader output = %q, want %q", got, want)
	}
}
