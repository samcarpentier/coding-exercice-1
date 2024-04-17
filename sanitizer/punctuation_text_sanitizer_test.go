package sanitizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TEXT_WITH_PUNCTUATION = "I'm: cool, am I; not!?"
)

func TestPunctuationTextSanitizer(t *testing.T) {
	t.Run("Given empty string when Sanitize then return empty string", func(t *testing.T) {
		// given
		sanitizer := &punctuationTextSanitizer{}

		// when
		sanitizedText := sanitizer.Sanitize("")

		// then
		assert.Equal(t, "", sanitizedText)
	})

	t.Run("Given string with punctuation when Sanitize then return string without punctuation but preserved apostrophes", func(t *testing.T) {
		// given
		sanitizer := &punctuationTextSanitizer{}

		// when
		sanitizedText := sanitizer.Sanitize(TEXT_WITH_PUNCTUATION)

		// then
		assert.Equal(t, "I'm cool am I not", sanitizedText)
	})
}
