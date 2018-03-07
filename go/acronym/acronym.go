// Package acronym contains a solution to the acronym Exercism exercise
package acronym

import (
	"strings"
)

// Abbreviate constructs an accronym based on capitalizing every word in the input string
func Abbreviate(s string) string {
	var words = strings.FieldsFunc(s, splitter)
	var abbreviation = ""
	for _, r := range words {
		abbreviation += strings.Title(string(r[0]))
	}
	return abbreviation
}

func splitter(r rune) bool {
	return r == ' ' || r == '-'
}
