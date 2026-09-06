/*
Package application run the application
*/
package application

import (
	"fmt"

	"github.com/PaesslerAG/jsonpath"
	"github.com/mattn/go-runewidth"
)

func analyseData(spec *Spec, dataArray []interface{}) {
	for _, item := range dataArray {
		for i, column := range spec.Columns {
			updateColumnWidth(spec, i, column, item)
		}
	}
}

func updateColumnWidth(spec *Spec, columnIndex int, column Column, item interface{}) {
	for _, path := range column.Path {
		value, err := jsonpath.Get(path, item)
		if err != nil {
			continue
		}
		if value == nil {
			continue
		}
		valueStr := fmt.Sprintf("%v", value)
		valueWidth := runewidth.StringWidth(valueStr)
		if valueWidth > column.Width {
			spec.Columns[columnIndex].Width = valueWidth
		}
		break
	}
}
