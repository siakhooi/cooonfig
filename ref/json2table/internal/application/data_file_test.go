package application

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadDataFromFile(t *testing.T) {
	tmpDir := t.TempDir()
	dataFile := filepath.Join(tmpDir, "data.json")
	want := []byte(`{"name":"alice"}`)
	if err := os.WriteFile(dataFile, want, 0o644); err != nil {
		t.Fatalf("failed to create data file: %v", err)
	}

	got, err := readData(dataFile)
	if err != nil {
		t.Fatalf("readData returned error: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("readData returned %q, want %q", string(got), string(want))
	}
}

func TestReadDataFromStdin(t *testing.T) {
	originalStdin := os.Stdin
	readPipe, writePipe, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdin pipe: %v", err)
	}
	os.Stdin = readPipe
	t.Cleanup(func() {
		os.Stdin = originalStdin
		_ = readPipe.Close()
	})

	want := `{"id":123}`
	go func() {
		_, _ = writePipe.Write([]byte(want))
		_ = writePipe.Close()
	}()

	got, err := readData("-")
	if err != nil {
		t.Fatalf("readData(-) returned error: %v", err)
	}
	if string(got) != want {
		t.Fatalf("readData(-) returned %q, want %q", string(got), want)
	}
}

func TestReadDataFileNotFound(t *testing.T) {
	_, err := readData(filepath.Join(t.TempDir(), "missing.json"))
	if err == nil {
		t.Fatal("expected error for missing file")
	}
	if !strings.Contains(err.Error(), "cannot read file:") {
		t.Fatalf("error = %q, want prefix containing %q", err.Error(), "cannot read file:")
	}
}

func TestReadDataReadFileError(t *testing.T) {
	dirPath := t.TempDir()

	_, err := readData(dirPath)
	if err == nil {
		t.Fatal("expected error when reading a directory as file")
	}
	if !strings.Contains(err.Error(), "error reading file:") {
		t.Fatalf("error = %q, want prefix containing %q", err.Error(), "error reading file:")
	}
}

func TestReadDataFromStdinReadError(t *testing.T) {
	originalStdin := os.Stdin
	tmpDir := t.TempDir()
	writeOnlyFile := filepath.Join(tmpDir, "write-only.txt")
	file, err := os.OpenFile(writeOnlyFile, os.O_CREATE|os.O_WRONLY, 0o600)
	if err != nil {
		t.Fatalf("failed to create write-only file: %v", err)
	}
	os.Stdin = file
	t.Cleanup(func() {
		os.Stdin = originalStdin
		_ = file.Close()
	})

	_, err = readData("-")
	if err == nil {
		t.Fatal("expected error when reading from invalid stdin")
	}
	if !strings.Contains(err.Error(), "error reading stdin:") {
		t.Fatalf("error = %q, want prefix containing %q", err.Error(), "error reading stdin:")
	}
}

func TestParseDataValidJSON(t *testing.T) {
	got, err := parseData([]byte(`{"name":"alice","age":30}`))
	if err != nil {
		t.Fatalf("parseData returned error: %v", err)
	}

	obj, ok := got.(map[string]interface{})
	if !ok {
		t.Fatalf("parseData returned %T, want map[string]interface{}", got)
	}
	if obj["name"] != "alice" {
		t.Fatalf("name = %v, want alice", obj["name"])
	}
	if obj["age"] != float64(30) {
		t.Fatalf("age = %v, want 30", obj["age"])
	}
}

func TestParseDataInvalidJSON(t *testing.T) {
	_, err := parseData([]byte(`{"name":`))
	if err == nil {
		t.Fatal("expected parseData to return error for invalid JSON")
	}
	if !strings.Contains(err.Error(), "error parsing JSON:") {
		t.Fatalf("error = %q, want prefix containing %q", err.Error(), "error parsing JSON:")
	}
}

func TestReadParseDataValid(t *testing.T) {
	tmpDir := t.TempDir()
	dataFile := filepath.Join(tmpDir, "data.json")
	if err := os.WriteFile(dataFile, []byte(`[{"id":1},{"id":2}]`), 0o644); err != nil {
		t.Fatalf("failed to create data file: %v", err)
	}

	got, err := readParseData(dataFile)
	if err != nil {
		t.Fatalf("ReadParseData returned error: %v", err)
	}

	arr, ok := got.([]interface{})
	if !ok {
		t.Fatalf("ReadParseData returned %T, want []interface{}", got)
	}
	if len(arr) != 2 {
		t.Fatalf("ReadParseData length = %d, want 2", len(arr))
	}
}

func TestReadParseDataInvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	dataFile := filepath.Join(tmpDir, "bad.json")
	if err := os.WriteFile(dataFile, []byte(`{"id":`), 0o644); err != nil {
		t.Fatalf("failed to create bad data file: %v", err)
	}

	_, err := readParseData(dataFile)
	if err == nil {
		t.Fatal("expected ReadParseData to return error for invalid JSON")
	}
	if !strings.Contains(err.Error(), "error parsing JSON:") {
		t.Fatalf("error = %q, want prefix containing %q", err.Error(), "error parsing JSON:")
	}
}
