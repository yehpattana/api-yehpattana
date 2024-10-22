package commonhelpers

import "strings"

func ReplacePercent20WithSpace(word string) string {
	// Remove leading "%20"
	for strings.HasPrefix(word, "%20") {
		word = strings.TrimPrefix(word, "%20")
	}
	// Remove trailing "%20"
	for strings.HasSuffix(word, "%20") {
		word = strings.TrimSuffix(word, "%20")
	}
	// Replace remaining "%20" with a space
	return strings.ReplaceAll(word, "%20", " ")
}
