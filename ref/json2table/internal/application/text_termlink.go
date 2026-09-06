/*
Package application run the application
*/
package application

import (
	"github.com/savioxavier/termlink"
)

func applyLink(text, url string) string {
	if url == "" {
		return text
	}
	return termlink.Link(text, url)
}
