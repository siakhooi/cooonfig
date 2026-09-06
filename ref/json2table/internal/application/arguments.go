/*
Package application run the application
*/
package application

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

// Arguments holds the parsed CLI arguments
type Arguments struct {
	SpecFile string
	DataFile string
	EnvSpec  string
	Columns  string
}

// ParseArguments parses CLI arguments, validates them into an Arguments struct
func ParseArguments(c *cli.Command) (Arguments, error) {
	envSpec := os.Getenv("JSON2TABLE_SPEC")

	columns := c.String("columns")
	specFile := validateSpecFile(c.String("spec"))

	argumentsList := c.Args().Slice()
	dataFile, err := validateDataFile(argumentsList)
	if err != nil {
		return Arguments{}, err
	}

	var errs []error
	if columns == "" {
		if specFile == "" && envSpec == "" {
			errs = append(errs, fmt.Errorf("spec is mandatory: provide -s/--spec flag or set JSON2TABLE_SPEC or JSON2TABLE_SPEC_FILE environment variable"))
		}
	}
	if len(argumentsList) > 1 {
		errs = append(errs, fmt.Errorf("too many arguments: only one data file argument is allowed"))
	}
	if dataFile == "" {
		errs = append(errs, fmt.Errorf("data file is required: provide a data file argument or pipe data to stdin"))
	}
	if len(errs) > 0 {
		return Arguments{}, errors.Join(errs...)
	}

	args := Arguments{
		SpecFile: specFile,
		DataFile: dataFile,
		EnvSpec:  envSpec,
		Columns:  columns,
	}

	return args, nil
}

func validateSpecFile(specFile string) string {
	// If specFile is provided via CLI flag, use it
	if specFile != "" {
		return specFile
	}

	// Try to read from environment variable
	envSpecFile := os.Getenv("JSON2TABLE_SPEC_FILE")
	if envSpecFile != "" {
		return envSpecFile
	}

	// Neither CLI flag nor environment variable provided
	return ""
}

func validateDataFile(args []string) (string, error) {
	if len(args) == 0 {
		// check if stdin is piped
		fi, err := os.Stdin.Stat()
		if err != nil {
			return "", fmt.Errorf("cannot stat stdin: %w", err)
		}
		// If stdin is not a character device, treat as piped input
		if (fi.Mode() & os.ModeCharDevice) == 0 {
			return "-", nil
		}

		return "", nil
	}

	filename := args[0]

	return filename, nil
}
