package sanitizer

// compositeTextSanitizer is a 'composite' that executes multiple 'TextSanitizer.Sanitize' functions and returns the result
type compositeTextSanitizer struct {
	Components []TextSanitizer
}

// newCompositeTextSanitizer private constructor that instantiates a new empty 'composite' TextSanitizer
func newCompositeTextSanitizer() *compositeTextSanitizer {
	return &compositeTextSanitizer{
		Components: []TextSanitizer{},
	}
}

// addTextSanitizer builder function that adds 1 sanitizer 'component' to the 'composite'
func (c *compositeTextSanitizer) addTextSanitizer(sanitizer TextSanitizer) {
	c.Components = append(c.Components, sanitizer)
}

// Sanitize executes the Sanitize function of every underlying 'components' in the 'composite' and returns the final result
func (c *compositeTextSanitizer) Sanitize(rawText string) string {
	var sanitizedText string = rawText

	for _, sanitizer := range c.Components {
		sanitizedText = sanitizer.Sanitize(sanitizedText)
	}

	return sanitizedText
}
