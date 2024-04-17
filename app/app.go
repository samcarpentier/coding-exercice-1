package app

import (
	"fmt"
	"log/slog"
	"os"

	"text/tabwriter"
	"trigrams/config"
	"trigrams/filereader"
	"trigrams/index"
	"trigrams/sanitizer"
)

// Run is the main execution of the program
func Run(config *config.Config) error {
	rawTextToParse, err := extractRawTextToParse(config)
	if err != nil {
		return err
	}

	sanitizer := sanitizer.NewTextSanitizer()
	sanitizedText := sanitizer.Sanitize(rawTextToParse)

	slog.Debug(sanitizedText)

	ngramindex := index.New(config.WordSequenceSize)
	ngramindex.CreateIndex(sanitizedText)
	topNWordsSequences := ngramindex.GetRankedSequencesByCount(config.NumberOfResultsToReturn)

	slog.Info("Results compilation completed!")
	printResultsAsTable(topNWordsSequences)

	return nil
}

// extractRawTextToParse returns the content of the files passed as input or the content of STDIN
func extractRawTextToParse(config *config.Config) (string, error) {
	if len(config.ExternalFilesList) > 0 {
		externalFilesReader := filereader.NewExternalFilesReader(config.ExternalFilesList)
		return externalFilesReader.ReadAllFilesContent()
	} else {
		return config.RawTextInput, nil
	}
}

// printResultsAsTable prints the top-n repeated words sequences as a human-readable table
func printResultsAsTable(topNWordsSequence []index.RepeatedWordsSequence) {
	tableWriter := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(tableWriter, "")
	fmt.Fprintln(tableWriter, "Words Sequence\tCount")
	fmt.Fprintln(tableWriter, "--------------------\t------")

	for _, wordsSequence := range topNWordsSequence {
		fmt.Fprintf(tableWriter, "%s\t%d\n", wordsSequence.WordsSequence, wordsSequence.Count)
	}

	tableWriter.Flush()
}
