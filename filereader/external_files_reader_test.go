package filereader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	FILE_PATH_1 = "./fixtures/simple_file1.txt"
	FILE_PATH_2 = "./fixtures/simple_file2.txt"
)

func TestExternalFilesReader(t *testing.T) {
	t.Run("Given single file to read when ReadAllFilesContent then file content is returned", func(t *testing.T) {
		// given
		reader := NewExternalFilesReader([]string{FILE_PATH_1})

		// when
		content, err := reader.ReadAllFilesContent()

		// then
		assert.NoError(t, err)
		assert.Equal(t, "I love mom's spaghetti\n", content)
	})

	t.Run("Given multiple files to read when ReadAllFilesContent then concatenated files content is returned", func(t *testing.T) {
		// given
		reader := NewExternalFilesReader([]string{FILE_PATH_1, FILE_PATH_2})

		// when
		content, err := reader.ReadAllFilesContent()

		// then
		assert.NoError(t, err)
		assert.Equal(t, "I love mom's spaghetti\nI like trains\n", content)
	})

	t.Run("Given file doesn't exist when ReadAllFilesContent then error is returned", func(t *testing.T) {
		// given
		reader := NewExternalFilesReader([]string{"UNKNOWN"})

		// when
		content, err := reader.ReadAllFilesContent()

		// then
		assert.Empty(t, content)
		assert.Error(t, err)
	})
}
