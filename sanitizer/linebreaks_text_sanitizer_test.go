package sanitizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TEXT_WITH_LINEBREAKS = "my\n  message"
)

func TestLineBreaksTextSanitizer(t *testing.T) {
	t.Run("Given empty string when Sanitize then return empty string", func(t *testing.T) {
		// given
		sanitizer := &lineBreaksTextSanitizer{}

		// when
		sanitizedText := sanitizer.Sanitize("")

		// then
		assert.Equal(t, "", sanitizedText)
	})

	t.Run("Given string with linebreaks and redundant spaces when Sanitize then return string without linebreaks nor redundant spaces", func(t *testing.T) {
		// given
		sanitizer := &lineBreaksTextSanitizer{}

		// when
		sanitizedText := sanitizer.Sanitize(TEXT_WITH_LINEBREAKS)

		// then
		assert.Equal(t, "my message", sanitizedText)
	})
}
