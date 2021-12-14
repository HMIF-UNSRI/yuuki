package helper

import "strings"

// ConvertNameToSlug Convert a name to slug:
// * Replace all whitespace with hyphen, lower the string to lowercase and
func ConvertNameToSlug(name string) string {
	// Split to words only
	words := strings.Fields(strings.ToLower(name))
	if len(words) == 0 {
		return ""
	}

	return strings.Join(words, "-")
}
