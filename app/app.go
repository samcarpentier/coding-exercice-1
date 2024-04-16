package app

import (
	"log/slog"
	"trigrams/config"
	"trigrams/parser"
)

// Run is the main execution of the program
func Run(config *config.Config) error {
	rawTextToParse, err := extractRawTextToParse(config)
	if err != nil {
		return err
	}

	sanitizer := parser.NewTextSanitizer()
	sanitizedText := sanitizer.Sanitize(rawTextToParse)

	slog.Info(sanitizedText)

	return nil
}

// extractRawTextToParse returns the content of the files passed as input or the content of STDIN
func extractRawTextToParse(config *config.Config) (string, error) {
	if len(config.ExternalFilesList) > 0 {
		externalFilesReader := parser.NewExternalFilesReader(config.ExternalFilesList)
		return externalFilesReader.ReadAllFilesContent()
	} else {
		return config.RawTextInput, nil
	}
}
