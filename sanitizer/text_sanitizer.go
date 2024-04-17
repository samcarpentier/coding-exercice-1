package sanitizer

// TextSanitizer represents the generic interface of a text sanitizer.
type TextSanitizer interface {
	Sanitize(string) string
}

// NewCompositeTextSanitizer public constructor that instantiates a fully configured 'composite' TextSanitizer
func NewTextSanitizer() TextSanitizer {
	compositeTextSanitizer := newCompositeTextSanitizer()
	compositeTextSanitizer.addTextSanitizer(&upperCaseTextSanitizer{})
	compositeTextSanitizer.addTextSanitizer(&punctuationTextSanitizer{})
	compositeTextSanitizer.addTextSanitizer(&lineBreaksTextSanitizer{})

	return compositeTextSanitizer
}
