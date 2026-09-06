/*
Package application run the application
*/
package application

func convertColumnsToSpec(columns string) (*Spec, error) {
	var spec Spec
	for _, col := range splitAndTrimCSV(columns) {
		spec.Columns = append(spec.Columns, Column{Path: StringList{col}})
	}
	spec.setDefaults()
	return &spec, nil
}

// splitAndTrimCSV splits a comma-separated string and trims spaces from each value
func splitAndTrimCSV(s string) []string {
	var result []string
	for _, v := range splitCSV(s) {
		trimmed := trimSpaces(v)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

// splitCSV splits a string by comma
func splitCSV(s string) []string {
	var res []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ',' {
			res = append(res, s[start:i])
			start = i + 1
		}
	}
	res = append(res, s[start:])
	return res
}

// trimSpaces trims leading and trailing spaces from a string
func trimSpaces(s string) string {
	start, end := 0, len(s)
	for start < end && (s[start] == ' ' || s[start] == '\t') {
		start++
	}
	for end > start && (s[end-1] == ' ' || s[end-1] == '\t') {
		end--
	}
	return s[start:end]
}
