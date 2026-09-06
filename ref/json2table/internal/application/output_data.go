/*
Package application run the application
*/
package application

import (
	"fmt"

	"github.com/PaesslerAG/jsonpath"
)

func selectFirstValue(paths StringList, item interface{}) interface{} {
	for _, path := range paths {
		value, err := jsonpath.Get(path, item)
		if err == nil && value != nil {
			return value
		}
	}

	return nil
}

func applyURLPath(printValue string, urlPath string, item interface{}) string {
	if urlPath == "" {
		return printValue
	}

	urlValue, err := jsonpath.Get(urlPath, item)
	if err != nil || urlValue == nil {
		return printValue
	}

	urlStr := fmt.Sprintf("%v", urlValue)
	return applyLink(printValue, urlStr)
}

func printData(dataArray []interface{}, spec *Spec) {
	for _, item := range dataArray {
		for _, column := range spec.Columns {
			value := selectFirstValue(column.Path, item)
			fullValue := fmt.Sprintf("%v", value)
			prefix, printValue, suffix := formatAlignedTextParts(fullValue, column.Width, column.Align)
			printValue = applyURLPath(printValue, column.URLPath, item)
			coloredPrintValue := applyColor(fullValue, printValue, column.Color)
			fmt.Printf("%s%s%s ", prefix, coloredPrintValue, suffix)
		}
		fmt.Println("")
	}

}
