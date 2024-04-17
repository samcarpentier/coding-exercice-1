package sanitizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	RAW_TEXT_INPUT = "Don't stop,\n   Believing!"
)

type testSanitizer struct {
	SanitizedWasCalled bool
}

func (s *testSanitizer) Sanitize(rawText string) string {
	s.SanitizedWasCalled = true
	return rawText
}

func TestCompositeTextSanitizer(t *testing.T) {
	t.Run("Given sanitizer components when AddTextSanitizer then sanitizers are added to the list", func(t *testing.T) {
		// given
		sanitizer := &compositeTextSanitizer{}

		// when
		sanitizer.addTextSanitizer(&testSanitizer{})
		sanitizer.addTextSanitizer(&testSanitizer{})
		sanitizer.addTextSanitizer(&testSanitizer{})

		// then
		assert.Len(t, sanitizer.Components, 3)
	})

	t.Run("Given no sanitizers in composite when Sanitize then original text is returned", func(t *testing.T) {
		// given
		sanitizer := &compositeTextSanitizer{}

		// when
		sanitizedText := sanitizer.Sanitize(RAW_TEXT_INPUT)

		// then
		assert.Equal(t, RAW_TEXT_INPUT, sanitizedText)
	})

	t.Run("Given empty string when Sanitize then empty string is returned", func(t *testing.T) {
		// given
		sanitizer := &compositeTextSanitizer{}

		// when
		sanitizedText := sanitizer.Sanitize("")

		// then
		assert.Equal(t, "", sanitizedText)
	})

	t.Run("Given composite with components when Sanitize then components Sanitize function is called", func(t *testing.T) {
		// given
		component1 := &testSanitizer{}
		component2 := &testSanitizer{}
		sanitizer := newCompositeTextSanitizer()
		sanitizer.addTextSanitizer(component1)
		sanitizer.addTextSanitizer(component2)

		// when
		sanitizer.Sanitize("anything")

		// then
		assert.True(t, component1.SanitizedWasCalled)
		assert.True(t, component2.SanitizedWasCalled)
	})
}
