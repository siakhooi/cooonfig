/*
Package application run the application
*/
package application

import (
	"encoding/json"
	"fmt"
)

// TextColorSpec represents text color specification
type TextColorSpec struct {
	Type       colorType
	Default    StringList `json:"default"`
	Conditions []struct {
		When  StringList `json:"when"`
		Color StringList `json:"color"`
	} `json:"conditions"`
}

// DefaultTextColor is the default color (no color)
var DefaultTextColor = TextColorSpec{
	Type: colorTypeFixed, Default: StringList{ColorDefault},
}

type colorType int

const (
	colorTypeFixed colorType = iota
	colorTypeConditional
)

// UnmarshalJSON implements json.Unmarshaler for TextColorSpec
func (s *TextColorSpec) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as a string or string array
	var fixed StringList
	err := json.Unmarshal(data, &fixed)
	if err == nil {
		*s = TextColorSpec{
			Type:    colorTypeFixed,
			Default: fixed,
		}
		return nil
	}

	// Try to unmarshal as an object
	var obj struct {
		Default    StringList `json:"default"`
		Conditions []struct {
			When  StringList `json:"when"`
			Color StringList `json:"color"`
		} `json:"conditions"`
	}
	err = json.Unmarshal(data, &obj)
	if err == nil {
		if obj.Default == nil {
			obj.Default = StringList{ColorDefault}
		}

		*s = TextColorSpec{
			Type:       colorTypeConditional,
			Default:    obj.Default,
			Conditions: obj.Conditions,
		}
		return nil
	}

	return fmt.Errorf("invalid color specification: %s", string(data))
}
