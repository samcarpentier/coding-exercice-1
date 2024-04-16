package config

import "errors"

// Config represents the configuration interface used within this application.
type Config struct {
	// NumberOfResultsToReturn represents the number of results to return.
	NumberOfResultsToReturn int

	// WordSequenceSize represents the size of the word sequences being gathered.
	WordSequenceSize int

	// ExternalFilesList contains a list of relative file paths to read to obtain the input text. Mutually exclusive with RawTextInput.
	ExternalFilesList []string

	// RawTextInput contains a sample of raw text to parse. Mutually exclusive with ExternalFilesList.
	RawTextInput string
}

// New inatantiates a Config structure.
func New(numberOfResultsToReturn int, wordSequenceSize int) *Config {
	return &Config{
		NumberOfResultsToReturn: numberOfResultsToReturn,
		WordSequenceSize:        wordSequenceSize,
	}
}

// AddExternalFile appends a file path to member attribute ExternalFilesList
func (c *Config) AddExternalFile(filePath string) error {
	if c.RawTextInput != "" {
		return errors.New("only one of ExternalFilesList or RawTextInput is supported")
	}

	c.ExternalFilesList = append(c.ExternalFilesList, filePath)
	return nil
}

// SetRawTextInput sets a string of raw text to member attribute RawTextInput
func (c *Config) SetRawTextInput(rawText string) error {
	if c.ExternalFilesList != nil {
		return errors.New("only one of ExternalFilesList or RawTextInput is supported")
	}

	c.RawTextInput = rawText
	return nil
}
