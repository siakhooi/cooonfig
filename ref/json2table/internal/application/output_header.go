/*
Package application run the application
*/
package application

import "fmt"

func printHeader(columns []Column) {
	for _, column := range columns {
		title := column.Title
		prefix, printValue, suffix := formatAlignedTextParts(fmt.Sprintf("%v", title), column.Width, column.Align)

		fmt.Printf("%s%s%s ", prefix, printValue, suffix)

	}
	fmt.Println("")
}
