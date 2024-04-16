package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	SEQUENCE_SIZE = 1

	SANITIZED_TEXT_SAMPLE = "W1 W2 W2 W3"
)

func TestNGramIndex(t *testing.T) {
	t.Run("Given sanitized text when CreateIndex then index is populated with word sequence occurrences count", func(t *testing.T) {
		// given
		idx := New(SEQUENCE_SIZE)

		// when
		idx.CreateIndex(SANITIZED_TEXT_SAMPLE)

		// then
		assert.Equal(t, 1, idx.Index["W1"])
		assert.Equal(t, 2, idx.Index["W2"])
		assert.Equal(t, 1, idx.Index["W3"])
	})

	t.Run("Given index is built when GetRankedSequencesByCount with 1 result then top-1 results are returned", func(t *testing.T) {
		// given
		idx := New(SEQUENCE_SIZE)
		idx.CreateIndex(SANITIZED_TEXT_SAMPLE)

		// when
		topOneRepeatedSequences := idx.GetRankedSequencesByCount(1)

		// then
		assert.Len(t, topOneRepeatedSequences, 1)
		assert.Equal(t, "W2", topOneRepeatedSequences[0].WordsSequence)
		assert.Equal(t, 2, topOneRepeatedSequences[0].Count)
	})

	t.Run("Given index is built when GetRankedSequencesByCount with multiple results then top-n results are returned", func(t *testing.T) {
		// given
		idx := New(SEQUENCE_SIZE)
		idx.CreateIndex(SANITIZED_TEXT_SAMPLE)

		// when
		topOneRepeatedSequences := idx.GetRankedSequencesByCount(3)

		// then
		assert.Len(t, topOneRepeatedSequences, 3)
		assert.Equal(t, "W2", topOneRepeatedSequences[0].WordsSequence)
		assert.Equal(t, 2, topOneRepeatedSequences[0].Count)
		assert.Equal(t, "W1", topOneRepeatedSequences[1].WordsSequence)
		assert.Equal(t, 1, topOneRepeatedSequences[1].Count)
		assert.Equal(t, "W3", topOneRepeatedSequences[2].WordsSequence)
		assert.Equal(t, 1, topOneRepeatedSequences[2].Count)
	})
}
