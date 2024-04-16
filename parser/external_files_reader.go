package parser

import (
	"os"
)

// ExternalFilesReader represents the module responsible of reading external files.
type ExternalFilesReader struct {
	// ExternalFilesList contains a list of relative file paths to read to obtain the input text. Mutually exclusive with RawTextInput.
	ExternalFilesList []string
}

// New instantiates a new ExternalFilesReader structure.
func NewExternalFilesReader(externalFilesList []string) *ExternalFilesReader {
	return &ExternalFilesReader{
		ExternalFilesList: externalFilesList,
	}
}

// ReadAllFilesContent reads the content of all files passed to the constructor
func (e *ExternalFilesReader) ReadAllFilesContent() (string, error) {
	var allFilesContent string
	for _, filename := range e.ExternalFilesList {
		content, err := e.readSingleFileContent(filename)
		if err != nil {
			return "", err
		}
		allFilesContent += content
	}

	return allFilesContent, nil
}

// readSingleFileContent reads the content of a single file
func (e *ExternalFilesReader) readSingleFileContent(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
