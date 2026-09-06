/*
Package application run the application
*/
package application

import "encoding/json"

// StringList is a custom type that can unmarshal from either a string or []string
type StringList []string

// UnmarshalJSON implements json.Unmarshaler for StringList
func (s *StringList) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as a string first
	if str := new(string); json.Unmarshal(data, str) == nil {
		*s = []string{*str}
		return nil
	}

	// Try to unmarshal as a string array
	var arr []string
	if err := json.Unmarshal(data, &arr); err != nil {
		return err
	}
	*s = arr
	return nil
}
