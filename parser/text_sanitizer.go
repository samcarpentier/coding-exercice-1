package parser

import (
	"regexp"
	"strings"
)

// TextSanitizer represents the module responsible of sanitizing the raw text input to standardize it and make it usable
type TextSanitizer struct {
}

func NewTextSanitizer() *TextSanitizer {
	return &TextSanitizer{}
}

// Sanitize takes raw text as input and returns a sanitized, ready to use string (case insensitive, no punctuation, no line breaks, no redundant spaces)
func (t *TextSanitizer) Sanitize(rawText string) string {
	var sanitizedText string

	sanitizedText = t.toUpperCase(rawText)
	sanitizedText = t.removePunctuationKeepContractions(sanitizedText)
	sanitizedText = t.removeLineBreaksAndRedundantSpaces(sanitizedText)

	return sanitizedText
}

// toUpper makes the whole raw text upper-case
func (t *TextSanitizer) toUpperCase(rawText string) string {
	return strings.ToUpper(rawText)
}

// removePunctuationKeepContractions filters only alphanumerical characters, apostrophes and spaces
func (t *TextSanitizer) removePunctuationKeepContractions(rawText string) string {
	return regexp.MustCompile(`[^'a-zA-Z0-9\s]+`).ReplaceAllString(rawText, "")
}

// removeBlankLinesAndRedundantSpaces replaces line breaks by spaces and strips all redundant spaces
func (t *TextSanitizer) removeLineBreaksAndRedundantSpaces(rawText string) string {
	var sanitizedText string

	sanitizedText = strings.ReplaceAll(rawText, "\n", " ")
	sanitizedText = strings.Join(strings.Fields(sanitizedText), " ")

	return sanitizedText
}
