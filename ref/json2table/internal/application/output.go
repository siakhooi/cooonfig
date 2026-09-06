/*
Package application run the application
*/
package application

// PrintTable prints JSON data in tabular format based on the provided specification
func PrintTable(spec *Spec, fullData interface{}) error {
	dataArray, err := selectDataArray(spec.DataPath, fullData)
	if err != nil {
		return err
	}
	analyseData(spec, dataArray)
	optimizeSpec(spec)

	printHeader(spec.Columns)
	printData(dataArray, spec)

	return nil
}
