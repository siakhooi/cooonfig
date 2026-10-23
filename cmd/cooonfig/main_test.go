package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const helperProcessEnv = "COOONFIG_TEST_MAIN"

func exampleDir(t *testing.T) string {
	t.Helper()
	dir, err := filepath.Abs(filepath.Join("..", "..", "example"))
	if err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestVersionString(t *testing.T) {
	if !strings.Contains(versionString(), Version) {
		t.Fatalf("versionString() = %q", versionString())
	}
}

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

func TestRunNoArgsProcesses(t *testing.T) {
	example := exampleDir(t)
	outputDir := t.TempDir()

	err := run([]string{
		"cooonfig",
		"-c", filepath.Join(example, "cooonfig.yaml"),
		"-i", example,
		"-o", outputDir,
	})
	if err != nil {
		t.Fatalf("process: %v", err)
	}

	for _, path := range []string{
		"apache/config.json",
		"kafka/config.yaml",
	} {
		if _, err := os.Stat(filepath.Join(outputDir, path)); err != nil {
			t.Fatalf("expected output %s: %v", path, err)
		}
	}
}

func TestRunNoArgsMissingConfigShowsHelp(t *testing.T) {
	t.Chdir(t.TempDir())

	output := captureStdout(t, func() {
		if err := run([]string{"cooonfig"}); err != nil {
			t.Fatalf("no args: %v", err)
		}
	})
	if !strings.Contains(output, "--help") || !strings.Contains(output, "--version") {
		t.Fatalf("expected help, got %q", output)
	}
}

func TestRunNoArgsExistingConfigProcesses(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "cooonfig.yaml"), nil, 0o644); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)

	var err error
	output := captureStdout(t, func() {
		err = run([]string{"cooonfig"})
	})
	if err == nil || !strings.Contains(err.Error(), "missing input files") {
		t.Fatalf("err = %v", err)
	}
	if strings.Contains(output, "--version") {
		t.Fatalf("expected processing, got help %q", output)
	}
}

func TestRunNoArgsStatError(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("directory mode is ignored for root")
	}

	dir := t.TempDir()
	t.Chdir(dir)
	t.Cleanup(func() {
		if err := os.Chmod(dir, 0o755); err != nil {
			t.Errorf("restore permissions: %v", err)
		}
	})
	if err := os.Chmod(dir, 0); err != nil {
		t.Fatal(err)
	}

	err := run([]string{"cooonfig"})
	if err == nil || !os.IsPermission(err) {
		t.Fatalf("err = %v, want permission error", err)
	}
}

func TestRunWithArgsShowsHelp(t *testing.T) {
	output := captureStdout(t, func() {
		if err := run([]string{"cooonfig", "anything"}); err != nil {
			t.Fatalf("args: %v", err)
		}
	})
	if !strings.Contains(output, "--help") || !strings.Contains(output, "--version") {
		t.Fatalf("expected help, got %q", output)
	}
}

func TestRunUnknownFlag(t *testing.T) {
	if err := run([]string{"cooonfig", "--spec"}); err == nil {
		t.Fatal("expected error for unknown flag")
	}
}

func TestMainSuccess(t *testing.T) {
	if os.Getenv(helperProcessEnv) == "1" {
		os.Args = []string{"cooonfig", "--help"}
		main()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=^TestMainSuccess$")
	cmd.Env = append(os.Environ(), helperProcessEnv+"=1")
	if err := cmd.Run(); err != nil {
		t.Fatalf("main success subprocess: %v", err)
	}
}

func TestMainErrorExit(t *testing.T) {
	if os.Getenv(helperProcessEnv) == "1" {
		os.Args = []string{"cooonfig", "--spec"}
		main()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=^TestMainErrorExit$")
	cmd.Env = append(os.Environ(), helperProcessEnv+"=1")
	err := cmd.Run()
	if err == nil {
		t.Fatal("expected subprocess exit error")
	}
	exitErr, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf("expected *exec.ExitError, got %T: %v", err, err)
	}
	if exitErr.ExitCode() != 1 {
		t.Fatalf("exit code = %d, want 1", exitErr.ExitCode())
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
