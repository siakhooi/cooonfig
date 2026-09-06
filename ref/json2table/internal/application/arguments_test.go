package application

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/urfave/cli/v3"
)

func runParseArguments(t *testing.T, argv []string) (Arguments, error) {
	t.Helper()

	var parsed Arguments
	cmd := &cli.Command{
		Name: "test",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "spec", Aliases: []string{"s"}},
		},
		Action: func(_ context.Context, c *cli.Command) error {
			var err error
			parsed, err = ParseArguments(c)
			return err
		},
	}

	err := cmd.Run(context.Background(), argv)
	return parsed, err
}

func TestValidateSpecFilePrefersCLIFlag(t *testing.T) {
	if err := os.Setenv("JSON2TABLE_SPEC_FILE", "env-spec.json"); err != nil {
		t.Fatalf("failed to set JSON2TABLE_SPEC_FILE: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Unsetenv("JSON2TABLE_SPEC_FILE")
	})

	got := validateSpecFile("cli-spec.json")
	if got != "cli-spec.json" {
		t.Fatalf("ValidateSpecFile returned %q, want %q", got, "cli-spec.json")
	}
}

func TestValidateSpecFileUsesEnvWhenCLIEmpty(t *testing.T) {
	if err := os.Setenv("JSON2TABLE_SPEC_FILE", "env-spec.json"); err != nil {
		t.Fatalf("failed to set JSON2TABLE_SPEC_FILE: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Unsetenv("JSON2TABLE_SPEC_FILE")
	})

	got := validateSpecFile("")
	if got != "env-spec.json" {
		t.Fatalf("ValidateSpecFile returned %q, want %q", got, "env-spec.json")
	}
}

func TestValidateSpecFileReturnsEmptyWhenMissing(t *testing.T) {
	if err := os.Unsetenv("JSON2TABLE_SPEC_FILE"); err != nil {
		t.Fatalf("failed to unset JSON2TABLE_SPEC_FILE: %v", err)
	}

	got := validateSpecFile("")
	if got != "" {
		t.Fatalf("ValidateSpecFile returned %q, want empty", got)
	}
}

func TestValidateDataFileWithArgument(t *testing.T) {
	got, err := validateDataFile([]string{"data.json"})
	if err != nil {
		t.Fatalf("ValidateDataFile returned error: %v", err)
	}
	if got != "data.json" {
		t.Fatalf("ValidateDataFile returned %q, want %q", got, "data.json")
	}
}

func TestValidateDataFileUsesStdinWhenPiped(t *testing.T) {
	originalStdin := os.Stdin
	readPipe, writePipe, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdin pipe: %v", err)
	}
	os.Stdin = readPipe
	t.Cleanup(func() {
		os.Stdin = originalStdin
		_ = readPipe.Close()
		_ = writePipe.Close()
	})

	got, err := validateDataFile(nil)
	if err != nil {
		t.Fatalf("ValidateDataFile returned error: %v", err)
	}
	if got != "-" {
		t.Fatalf("ValidateDataFile returned %q, want %q", got, "-")
	}
}

func TestValidateDataFileNoArgAndNoPipeReturnsEmpty(t *testing.T) {
	originalStdin := os.Stdin
	devNull, err := os.Open("/dev/null")
	if err != nil {
		t.Fatalf("failed to open /dev/null: %v", err)
	}
	os.Stdin = devNull
	t.Cleanup(func() {
		os.Stdin = originalStdin
		_ = devNull.Close()
	})

	got, err := validateDataFile(nil)
	if err != nil {
		t.Fatalf("ValidateDataFile returned error: %v", err)
	}
	if got != "" {
		t.Fatalf("ValidateDataFile returned %q, want empty", got)
	}
}

func TestValidateDataFileReturnsErrorWhenStdinStatFails(t *testing.T) {
	originalStdin := os.Stdin
	devNull, err := os.Open("/dev/null")
	if err != nil {
		t.Fatalf("failed to open /dev/null: %v", err)
	}
	if err := devNull.Close(); err != nil {
		t.Fatalf("failed to close /dev/null: %v", err)
	}
	os.Stdin = devNull
	t.Cleanup(func() {
		os.Stdin = originalStdin
	})

	_, err = validateDataFile(nil)
	if err == nil {
		t.Fatal("expected ValidateDataFile to return error")
	}
	if !strings.Contains(err.Error(), "cannot stat stdin") {
		t.Fatalf("error = %q, want to contain %q", err.Error(), "cannot stat stdin")
	}
}

func TestParseArgumentsWithSpecFlagAndDataFile(t *testing.T) {
	if err := os.Unsetenv("JSON2TABLE_SPEC"); err != nil {
		t.Fatalf("failed to unset JSON2TABLE_SPEC: %v", err)
	}
	if err := os.Unsetenv("JSON2TABLE_SPEC_FILE"); err != nil {
		t.Fatalf("failed to unset JSON2TABLE_SPEC_FILE: %v", err)
	}

	args, err := runParseArguments(t, []string{"test", "--spec", "spec.json", "data.json"})
	if err != nil {
		t.Fatalf("ParseArguments returned error: %v", err)
	}
	if args.SpecFile != "spec.json" {
		t.Fatalf("SpecFile = %q, want %q", args.SpecFile, "spec.json")
	}
	if args.DataFile != "data.json" {
		t.Fatalf("DataFile = %q, want %q", args.DataFile, "data.json")
	}
	if args.EnvSpec != "" {
		t.Fatalf("EnvSpec = %q, want empty", args.EnvSpec)
	}
}

func TestParseArgumentsWithEnvSpecAndPipedData(t *testing.T) {
	if err := os.Setenv("JSON2TABLE_SPEC", `{"columns":[]}`); err != nil {
		t.Fatalf("failed to set JSON2TABLE_SPEC: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Unsetenv("JSON2TABLE_SPEC")
	})

	originalStdin := os.Stdin
	readPipe, writePipe, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdin pipe: %v", err)
	}
	os.Stdin = readPipe
	t.Cleanup(func() {
		os.Stdin = originalStdin
		_ = readPipe.Close()
		_ = writePipe.Close()
	})

	args, err := runParseArguments(t, []string{"test"})
	if err != nil {
		t.Fatalf("ParseArguments returned error: %v", err)
	}
	if args.SpecFile != "" {
		t.Fatalf("SpecFile = %q, want empty", args.SpecFile)
	}
	if args.DataFile != "-" {
		t.Fatalf("DataFile = %q, want %q", args.DataFile, "-")
	}
	if args.EnvSpec == "" {
		t.Fatal("EnvSpec is empty, want non-empty")
	}
}

func TestParseArgumentsReturnsJoinedErrors(t *testing.T) {
	if err := os.Unsetenv("JSON2TABLE_SPEC"); err != nil {
		t.Fatalf("failed to unset JSON2TABLE_SPEC: %v", err)
	}
	if err := os.Unsetenv("JSON2TABLE_SPEC_FILE"); err != nil {
		t.Fatalf("failed to unset JSON2TABLE_SPEC_FILE: %v", err)
	}

	originalStdin := os.Stdin
	devNull, err := os.Open("/dev/null")
	if err != nil {
		t.Fatalf("failed to open /dev/null: %v", err)
	}
	os.Stdin = devNull
	t.Cleanup(func() {
		os.Stdin = originalStdin
		_ = devNull.Close()
	})

	_, err = runParseArguments(t, []string{"test"})
	if err == nil {
		t.Fatal("expected ParseArguments to return error")
	}

	errText := err.Error()
	if !strings.Contains(errText, "spec is mandatory") {
		t.Fatalf("error = %q, want to contain %q", errText, "spec is mandatory")
	}
	if !strings.Contains(errText, "data file is required") {
		t.Fatalf("error = %q, want to contain %q", errText, "data file is required")
	}
}

func TestParseArgumentsTooManyArguments(t *testing.T) {
	if err := os.Setenv("JSON2TABLE_SPEC", `{"columns":[]}`); err != nil {
		t.Fatalf("failed to set JSON2TABLE_SPEC: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Unsetenv("JSON2TABLE_SPEC")
	})

	_, err := runParseArguments(t, []string{"test", "a.json", "b.json"})
	if err == nil {
		t.Fatal("expected ParseArguments to return error")
	}
	if !strings.Contains(err.Error(), "too many arguments") {
		t.Fatalf("error = %q, want to contain %q", err.Error(), "too many arguments")
	}
}

func TestParseArgumentsReturnsDataFileValidationError(t *testing.T) {
	if err := os.Setenv("JSON2TABLE_SPEC", `{"columns":[]}`); err != nil {
		t.Fatalf("failed to set JSON2TABLE_SPEC: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Unsetenv("JSON2TABLE_SPEC")
	})

	originalStdin := os.Stdin
	devNull, err := os.Open("/dev/null")
	if err != nil {
		t.Fatalf("failed to open /dev/null: %v", err)
	}
	if err := devNull.Close(); err != nil {
		t.Fatalf("failed to close /dev/null: %v", err)
	}
	os.Stdin = devNull
	t.Cleanup(func() {
		os.Stdin = originalStdin
	})

	_, err = runParseArguments(t, []string{"test"})
	if err == nil {
		t.Fatal("expected ParseArguments to return error")
	}
	if !strings.Contains(err.Error(), "cannot stat stdin") {
		t.Fatalf("error = %q, want to contain %q", err.Error(), "cannot stat stdin")
	}
}
