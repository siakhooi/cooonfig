/*
Package application run the application
*/
package application

import (
	"math"
	"strings"
)

// Column represents a column specification
type Column struct {
	Path     StringList    `json:"path" validate:"required"`
	Title    string        `json:"title"`
	MinWidth int           `json:"minWidth" validate:"min=0,ltefield=MaxWidth"`
	MaxWidth int           `json:"maxWidth" validate:"min=0,gtefield=MinWidth"`
	Align    Alignment     `json:"align" validate:"omitempty,oneof=left right center"`
	URLPath  string        `json:"urlPath"`
	Color    TextColorSpec `json:"color"`

	Width int
}

// Spec represents the specification structure
type Spec struct {
	DataPath string   `json:"dataPath"`
	Columns  []Column `json:"columns" validate:"required,min=1,dive"`
}

func (c *Column) setDefaults() {
	if c.Title == "" {
		parts := strings.Split(c.Path[0], ".")
		if len(parts) > 1 {
			c.Title = parts[len(parts)-1]
		} else {
			c.Title = c.Path[0]
		}
	}
	c.Width = len(c.Title)
	if c.MaxWidth == 0 {
		c.MaxWidth = math.MaxInt
	}
	if c.Align == "" {
		c.Align = DefaultAlignment
	}
	if len(c.Color.Default) == 0 {
		c.Color = DefaultTextColor
	}
}
func (s *Spec) setDefaults() {
	if s.DataPath == "" {
		s.DataPath = "$"
	}
	for i := range s.Columns {
		s.Columns[i].setDefaults()
	}
}
