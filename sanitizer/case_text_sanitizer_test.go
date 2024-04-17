package sanitizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TEXT_WITH_LOWERCASE = "my message"
)

func TestUpperCaseTextSanitizer(t *testing.T) {
	t.Run("Given empty string when Sanitize then return empty string", func(t *testing.T) {
		// given
		sanitizer := &upperCaseTextSanitizer{}

		// when
		sanitizedText := sanitizer.Sanitize("")

		// then
		assert.Equal(t, "", sanitizedText)
	})

	t.Run("Given string with lower case letters when Sanitize then string is capitalized", func(t *testing.T) {
		// given
		sanitizer := &upperCaseTextSanitizer{}

		// when
		sanitizedText := sanitizer.Sanitize(TEXT_WITH_LOWERCASE)

		// then
		assert.Equal(t, "MY MESSAGE", sanitizedText)
	})
}
