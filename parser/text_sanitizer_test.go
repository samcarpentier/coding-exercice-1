package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	RAW_TEXT_INPUT = "I, don't have    anything to say...\nDo you?"
)

func TestTextSanitizer(t *testing.T) {
	t.Run("given string when Sanitize then return fully sanitized text", func(t *testing.T) {
		// given
		sanitizer := NewTextSanitizer()

		// when
		sanitizedText := sanitizer.Sanitize(RAW_TEXT_INPUT)

		// then
		assert.Equal(t, "I DON'T HAVE ANYTHING TO SAY DO YOU", sanitizedText)
	})
}
