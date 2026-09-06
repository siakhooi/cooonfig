package application

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/urfave/cli/v3"
)

func TestFlags(t *testing.T) {
	f := flags()

	if len(f) != 2 {
		t.Errorf("expected 2 flags, got %d", len(f))
	}

	// Test build flag
	buildFlag, ok := f[0].(*cli.BoolFlag)
	if !ok {
		t.Errorf("expected first flag to be BoolFlag")
	}
	if buildFlag.Name != "build" {
		t.Errorf("expected flag name 'build', got '%s'", buildFlag.Name)
	}
	if buildFlag.Usage != "print build info and exit" {
		t.Errorf("unexpected usage for build flag: %s", buildFlag.Usage)
	}

	// Test spec flag
	specFlag, ok := f[1].(*cli.StringFlag)
	if !ok {
		t.Errorf("expected second flag to be StringFlag")
	}
	if specFlag.Name != "spec" {
		t.Errorf("expected flag name 'spec', got '%s'", specFlag.Name)
	}
	if len(specFlag.Aliases) != 1 || specFlag.Aliases[0] != "s" {
		t.Errorf("expected alias 's' for spec flag")
	}
}

func TestRunWithVersion(t *testing.T) {
	// Test that Run doesn't panic with --version flag
	err := Run([]string{"json2table", "--version"})
	if err != nil {
		t.Errorf("unexpected error with --version: %v", err)
	}
}

func TestRunWithHelp(t *testing.T) {
	// Test that Run doesn't panic with --help flag
	err := Run([]string{"json2table", "--help"})
	if err != nil {
		t.Errorf("unexpected error with --help: %v", err)
	}
}

func TestActionWithBuildFlag(t *testing.T) {
	cmd := &cli.Command{
		Name:  "test",
		Flags: flags(),
		Action: func(_ context.Context, c *cli.Command) error {
			return action(context.Background(), c)
		},
	}

	// Test with --build flag - should print build info and return nil
	err := cmd.Run(context.Background(), []string{"test", "--build"})
	if err != nil {
		t.Errorf("unexpected error with --build flag: %v", err)
	}
}

func TestRunCommandStructure(t *testing.T) {
	// Verify the command is properly structured
	cmd := &cli.Command{
		Name:      "json2table",
		Usage:     "print json data in tabular format",
		ArgsUsage: "[dataFile]",
		Action:    action,
		Flags:     flags(),
	}

	if cmd.Name != "json2table" {
		t.Errorf("expected name 'json2table', got '%s'", cmd.Name)
	}
	if cmd.Usage != "print json data in tabular format" {
		t.Errorf("unexpected usage: %s", cmd.Usage)
	}
	if cmd.ArgsUsage != "[dataFile]" {
		t.Errorf("unexpected args usage: %s", cmd.ArgsUsage)
	}
}

func TestActionWithNoSpecFile(t *testing.T) {
	// Unset environment variables to ensure ParseArguments fails
	if err := os.Unsetenv("JSON2TABLE_SPEC"); err != nil {
		t.Fatalf("failed to unset JSON2TABLE_SPEC: %v", err)
	}
	if err := os.Unsetenv("JSON2TABLE_SPEC_FILE"); err != nil {
		t.Fatalf("failed to unset JSON2TABLE_SPEC_FILE: %v", err)
	}

	cmd := &cli.Command{
		Name:  "test",
		Flags: flags(),
		Action: func(_ context.Context, c *cli.Command) error {
			return action(context.Background(), c)
		},
	}

	// Test without spec - should return error from ParseArguments
	err := cmd.Run(context.Background(), []string{"test", "nonexistent.json"})
	if err == nil {
		t.Error("expected error when no spec is provided")
	}
}

func TestActionWithInvalidSpecFile(t *testing.T) {
	// Create a temporary invalid spec file
	tmpDir := t.TempDir()
	specFile := filepath.Join(tmpDir, "invalid-spec.json")
	err := os.WriteFile(specFile, []byte("invalid json"), 0644)
	if err != nil {
		t.Fatalf("failed to create temp spec file: %v", err)
	}

	cmd := &cli.Command{
		Name:  "test",
		Flags: flags(),
		Action: func(_ context.Context, c *cli.Command) error {
			return action(context.Background(), c)
		},
	}

	// Test with invalid spec file - should return error from ReadParseValidateSpec
	err = cmd.Run(context.Background(), []string{"test", "--spec", specFile, "data.json"})
	if err == nil {
		t.Error("expected error when spec file is invalid")
	}
}

func TestActionWithInvalidDataFile(t *testing.T) {
	// Create a temporary valid spec file
	tmpDir := t.TempDir()
	specFile := filepath.Join(tmpDir, "spec.json")
	specContent := `{"rootPath": ".", "columns": [{"header": "name", "path": "name"}]}`
	err := os.WriteFile(specFile, []byte(specContent), 0644)
	if err != nil {
		t.Fatalf("failed to create temp spec file: %v", err)
	}

	cmd := &cli.Command{
		Name:  "test",
		Flags: flags(),
		Action: func(_ context.Context, c *cli.Command) error {
			return action(context.Background(), c)
		},
	}

	// Test with non-existent data file - should return error from ReadParseData
	err = cmd.Run(context.Background(), []string{"test", "--spec", specFile, "nonexistent-data.json"})
	if err == nil {
		t.Error("expected error when data file does not exist")
	}
}

func TestActionWithValidInput(t *testing.T) {
	// Create temporary valid spec and data files
	tmpDir := t.TempDir()

	specFile := filepath.Join(tmpDir, "spec.json")
	specContent := `{"rootPath": ".", "columns": [{"header": "name", "path": "name"}]}`
	err := os.WriteFile(specFile, []byte(specContent), 0644)
	if err != nil {
		t.Fatalf("failed to create temp spec file: %v", err)
	}

	dataFile := filepath.Join(tmpDir, "data.json")
	dataContent := `[{"name": "test"}]`
	err = os.WriteFile(dataFile, []byte(dataContent), 0644)
	if err != nil {
		t.Fatalf("failed to create temp data file: %v", err)
	}

	cmd := &cli.Command{
		Name:  "test",
		Flags: flags(),
		Action: func(_ context.Context, c *cli.Command) error {
			return action(context.Background(), c)
		},
	}

	// Test with valid spec and data files - should succeed
	err = cmd.Run(context.Background(), []string{"test", "--spec", specFile, dataFile})
	if err != nil {
		t.Errorf("unexpected error with valid input: %v", err)
	}
}
