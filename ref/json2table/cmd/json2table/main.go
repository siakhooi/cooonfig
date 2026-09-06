/*
main cli entry
*/
package main

import (
	"fmt"
	"os"

	"github.com/siakhooi/json2table/internal/application"
)

func main() {

	if err := application.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
