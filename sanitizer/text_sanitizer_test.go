package sanitizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextSanitizer(t *testing.T) {
	t.Run("Given sanitizer instantiated with public constructor when Sanitize then return fully sanitized string", func(t *testing.T) {
		// given
		sanitizer := NewTextSanitizer()

		// when
		sanitizedText := sanitizer.Sanitize(RAW_TEXT_INPUT)

		// then
		assert.Equal(t, "DON'T STOP BELIEVING", sanitizedText)
	})
}
