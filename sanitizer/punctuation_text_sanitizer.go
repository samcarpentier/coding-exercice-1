package sanitizer

import (
	"regexp"
)

// punctuationTextSanitizer is a 'component' sanitizer that removes punctuation from the text.
type punctuationTextSanitizer struct {
}

// Sanitize executes the sanitization for 'component' punctuationTextSanitizer.
func (u *punctuationTextSanitizer) Sanitize(rawText string) string {
	return regexp.MustCompile(`[^'a-zA-Z0-9\s]+`).ReplaceAllString(rawText, "")
}
