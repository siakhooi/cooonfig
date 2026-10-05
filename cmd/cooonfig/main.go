// Package main is the CLI entrypoint.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

var (
	// Version is the release version. Override with -ldflags.
	Version = "0.0.0"
	// Date is the UTC build time. Override with -ldflags.
	Date = "unknown"
	// Commit is the source revision. Override with -ldflags.
	Commit = "unknown"
)

func versionString() string {
	return fmt.Sprintf("%s (commit %s, built %s)", Version, Commit, Date)
}

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	cmd := &cli.Command{
		Name:    "cooonfig",
		Usage:   "transform config files",
		Version: versionString(),
		Action:  action,
	}
	return cmd.Run(context.Background(), args)
}

func action(_ context.Context, c *cli.Command) error {
	if c.Args().Len() == 0 {
		return cli.ShowAppHelp(c)
	}
	return nil
}
