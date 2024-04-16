package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	NUM_RESULTS   = 10
	SEQUENCE_SIZE = 2

	EXTERNAL_FILE_PATH_1 = "./texts/file1.txt"
	EXTERNAL_FILE_PATH_2 = "./texts/file2.txt"

	RAW_TEXT = "Lorem ipsum dolor sit amet"
)

func TestConfig(t *testing.T) {
	t.Run("Given valid config struct when AddExternalFile then files are appended to list", func(t *testing.T) {
		// given
		cfg := New(NUM_RESULTS, SEQUENCE_SIZE)

		// when
		cfg.AddExternalFile(EXTERNAL_FILE_PATH_1)
		cfg.AddExternalFile(EXTERNAL_FILE_PATH_2)

		// then
		assert.Len(t, cfg.ExternalFilesList, 2)
		assert.Contains(t, cfg.ExternalFilesList, EXTERNAL_FILE_PATH_1)
		assert.Contains(t, cfg.ExternalFilesList, EXTERNAL_FILE_PATH_2)
	})

	t.Run("Given valid config struct when SetRawTextInput then raw text input is set", func(t *testing.T) {
		// given
		cfg := New(NUM_RESULTS, SEQUENCE_SIZE)

		// when
		cfg.SetRawTextInput(RAW_TEXT)

		// then
		assert.Equal(t, cfg.RawTextInput, RAW_TEXT)
	})

	t.Run("Given config with raw text input when AddExternalFile then error is returned", func(t *testing.T) {
		// given
		cfg := New(NUM_RESULTS, SEQUENCE_SIZE)
		cfg.SetRawTextInput(RAW_TEXT)

		// when
		err := cfg.AddExternalFile(EXTERNAL_FILE_PATH_1)

		// then
		assert.Error(t, err)
	})

	t.Run("Given config with external files in list when SetRawTextInput then error is returned", func(t *testing.T) {
		// given
		cfg := New(NUM_RESULTS, SEQUENCE_SIZE)
		cfg.AddExternalFile(EXTERNAL_FILE_PATH_1)

		// when
		err := cfg.SetRawTextInput(RAW_TEXT)

		// then
		assert.Error(t, err)
	})
}
