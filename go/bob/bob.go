//Package bob contains a solution to the bob exercism exercise
package bob

import (
	"strings"
	"unicode"
)

// Hey inspects the input string (remark), and returns the appropriate response
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	var response string
	switch {
	case remark == "":
		response = "Fine. Be that way!"
	case strings.HasSuffix(remark, "?") && !isShouting(remark):
		response = "Sure."
	case strings.HasSuffix(remark, "?") && isShouting(remark):
		response = "Calm down, I know what I'm doing!"
	case isShouting(remark):
		response = "Whoa, chill out!"
	default:
		response = "Whatever."
	}
	return response
}

func isShouting(remark string) bool {
	var testString = strings.ToUpper(remark)
	if testString == remark && containsLetter(remark) {
		return true
	}
	return false
}

func containsLetter(remark string) bool {
	for _, r := range remark {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
