package index

import (
	"fmt"
	"log/slog"
	"sort"
	"strings"
)

// NGramIndex represents the n-gram index.
type NGramIndex struct {
	// SequenceSize represents the number of words to consider in the sequence (with s=3, we are parsing trigrams).
	SequenceSize int

	// Index is a map for which the key is the n-words sequence and the value is the number of occurrences of this sequence in the input text.
	Index map[string]int
}

// RepeatedWordSequence represents a repeated word sequence in the input text.
type RepeatedWordsSequence struct {
	// WordsSequence is the n-words sequence.
	WordsSequence string

	// Count is the number of times this sequence is repeated in the text.
	Count int
}

// New instantiates a new NGramIndex structure.
func New(sequenceSize int) *NGramIndex {
	return &NGramIndex{
		SequenceSize: sequenceSize,
		Index:        make(map[string]int),
	}
}

// CreateIndex takes the sanitized text as input, counts and indexes the n-words sequences and store the result in the Index attribute of the NGramIndex structure.
func (n *NGramIndex) CreateIndex(sanitizedTextInput string) {
	inputTextWords := strings.Fields(sanitizedTextInput)

	for i := range inputTextWords {
		if i+n.SequenceSize <= len(inputTextWords) {
			wordsSequence := strings.Join(inputTextWords[i:(i+n.SequenceSize)], " ")
			_, existingWordsSequence := n.Index[wordsSequence]
			if existingWordsSequence {
				slog.Debug("Increment counter for existing sequence in index", "seq", wordsSequence, "newCount", n.Index[wordsSequence]+1)
				n.Index[wordsSequence] += 1
			} else {
				slog.Debug("Add new word sequence to index", "seq", wordsSequence)
				n.Index[wordsSequence] = 1
			}
		}
	}
}

// GetRankedSequencesByCount uses the index built by CreateIndex and returns the top-n reoccurring words sequences.
func (n *NGramIndex) GetRankedSequencesByCount(numberOfResults int) []RepeatedWordsSequence {
	slog.Info(fmt.Sprintf("Calculating the top-%d repeated %d word sequences...", numberOfResults, n.SequenceSize))

	var repeatedWordsSequences []RepeatedWordsSequence
	for wordSequence, count := range n.Index {
		repeatedWordsSequences = append(repeatedWordsSequences, RepeatedWordsSequence{wordSequence, count})
	}

	sort.Slice(repeatedWordsSequences, func(i int, j int) bool {
		return repeatedWordsSequences[i].Count > repeatedWordsSequences[j].Count
	})

	return repeatedWordsSequences[:numberOfResults]
}
