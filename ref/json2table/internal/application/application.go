/*
Package application run the application
*/
package application

import (
	"context"

	"github.com/siakhooi/json2table/internal/versioninfo"
	"github.com/urfave/cli/v3"
)

/*
Run start application
*/
func Run(args []string) error {
	cmd := &cli.Command{
		Name:      "json2table",
		Usage:     "print json data in tabular format",
		ArgsUsage: "[dataFile]",
		Version:   versioninfo.Version,
		Action:    action,
		Flags:     flags(),
	}
	return cmd.Run(context.Background(), args)

}

func action(_ context.Context, c *cli.Command) error {
	if c.Bool("build") {
		versioninfo.PrintBuildInfo()
		return nil
	}

	args, err := ParseArguments(c)
	if err != nil {
		return err
	}

	spec, err := readParseValidateSpec(args.SpecFile, args.EnvSpec, args.Columns)
	if err != nil {
		return err
	}

	jsonData, err := readParseData(args.DataFile)
	if err != nil {
		return err
	}
	return PrintTable(spec, jsonData)
}
func flags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "build",
			Usage: "print build info and exit",
		},
		&cli.StringFlag{
			Name:    "spec",
			Aliases: []string{"s"},
			Usage:   "read spec from specFile.json, or from environment variable JSON2TABLE_SPEC or JSON2TABLE_SPEC_FILE if not provided",
		},
		&cli.StringFlag{
			Name:    "columns",
			Aliases: []string{"c"},
			Usage:   "Comma separated list of columns to print, ignore -s and JSON2TABLE_SPEC or JSON2TABLE_SPEC_FILE if provided",
		},
	}
}
