package sanitizer

import (
	"strings"
)

// lineBreaksTextSanitizer is a 'component' sanitizer that removes line breaks and redundant spaces.
type lineBreaksTextSanitizer struct {
}

// Sanitize executes the sanitization for 'component' lineBreaksTextSanitizer.
func (u *lineBreaksTextSanitizer) Sanitize(rawText string) string {
	var sanitizedText string

	sanitizedText = strings.ReplaceAll(rawText, "\n", " ")
	sanitizedText = strings.Join(strings.Fields(sanitizedText), " ")

	return sanitizedText
}
