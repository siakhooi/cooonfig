/*
Package application run the application
*/
package application

import (
	"fmt"

	"github.com/PaesslerAG/jsonpath"
)

func selectDataArray(dataPath string, fullData interface{}) ([]interface{}, error) {
	data, err := jsonpath.Get(dataPath, fullData)
	if err != nil {
		return nil, fmt.Errorf("error selecting data with jsonpath: %w", err)
	}

	if dataArray, ok := data.([]interface{}); ok {
		return dataArray, nil
	}

	return nil, fmt.Errorf("data selected with jsonpath is not an array")
}
