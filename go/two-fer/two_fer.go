// Package twofer contains a solution to the gigasecond exercism exercise
package twofer

// ShareWith inspects the input string, and returns the appropriate twofer string
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
