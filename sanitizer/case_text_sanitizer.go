package sanitizer

import (
	"strings"
)

// upperCaseTextSanitizer is a 'component' sanitizer that makes the text upper case.
type upperCaseTextSanitizer struct {
}

// Sanitize executes the sanitization for 'component' upperCaseTextSanitizer.
func (u *upperCaseTextSanitizer) Sanitize(rawText string) string {
	return strings.ToUpper(rawText)
}
