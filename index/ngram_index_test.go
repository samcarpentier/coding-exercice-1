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
		topThreeRepeatedSequences := idx.GetRankedSequencesByCount(3)

		// then
		assert.Len(t, topThreeRepeatedSequences, 3)
		assert.Equal(t, "W2", topThreeRepeatedSequences[0].WordsSequence)
		assert.Equal(t, 2, topThreeRepeatedSequences[0].Count)

		// Since W1 and W3 are both at count 1, I used assert.ElementsMatch which disregards the order
		// in the slice. If I attempted to match the indexes 1 and 2 of the slice exactly, I would end
		// up with a "flaky" test because the order of same-count sequences is not deterministic with
		// my current implementation.
		expectedTopRepeatedSequences := []RepeatedWordsSequence{
			{WordsSequence: "W2", Count: 2},
			{WordsSequence: "W1", Count: 1},
			{WordsSequence: "W3", Count: 1},
		}
		assert.ElementsMatch(t, expectedTopRepeatedSequences, topThreeRepeatedSequences)
	})
}
