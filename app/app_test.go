package app

import (
	"io"
	"os"
	"trigrams/config"

	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	NUMBER_OF_RESULTS = 4
	SEQUENCE_SIZE     = 1

	FILE_PATH = "./fixtures/sample_text.txt"
	RAW_TEXT  = `
w1 w2 w3 w4
w1 w2 w3
w1 w2
w1
`
	EXPECTED_RESULTS_TABLE = `
Words Sequence       Count
-------------------- ------
W1                   4
W2                   3
W3                   2
W4                   1
`
)

func TestApp(t *testing.T) {
	t.Run("[END-TO-END] Given text input from file(s) when Run then results are printed", func(t *testing.T) {
		// given
		cfg := config.New(NUMBER_OF_RESULTS, SEQUENCE_SIZE)
		cfg.AddExternalFile(FILE_PATH)

		// when
		var err error
		textOutput := captureOutput(func() {
			err = Run(cfg)
		})

		// then
		assert.NoError(t, err)
		assert.Equal(t, EXPECTED_RESULTS_TABLE, textOutput)
	})

	t.Run("[END-TO-END] Given raw text input when Run then results are printed", func(t *testing.T) {
		// given
		cfg := config.New(NUMBER_OF_RESULTS, SEQUENCE_SIZE)
		cfg.SetRawTextInput(RAW_TEXT)

		// when
		var err error
		textOutput := captureOutput(func() {
			err = Run(cfg)
		})

		// then
		assert.NoError(t, err)
		assert.Equal(t, EXPECTED_RESULTS_TABLE, textOutput)
	})
}

func captureOutput(f func()) string {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	os.Stdout = orig
	w.Close()
	out, _ := io.ReadAll(r)
	return string(out)
}
