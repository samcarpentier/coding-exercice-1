package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"

	"trigrams/app"
	"trigrams/config"
)

const (
	DEFAULT_NUM_RESULTS_TO_RETURN = 100
	DEFAULT_WORD_SEQUENCE_SIZE    = 3
)

var (
	numberOfResultsToReturn int
	wordSequenceSize        int
	isVerbose               bool
)

func main() {
	// Parse CLI arguments
	flag.IntVar(&numberOfResultsToReturn, "n", DEFAULT_NUM_RESULTS_TO_RETURN, "Number of results to return")
	flag.IntVar(&wordSequenceSize, "s", DEFAULT_WORD_SEQUENCE_SIZE, "The size of the word sequence to capture")
	flag.BoolVar(&isVerbose, "v", false, "Enables verbose logging")
	flag.Parse()

	// Configure global logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if isVerbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	} else {
		slog.SetLogLoggerLevel(slog.LevelInfo)
	}

	// Build configuration struct
	cfg := config.New(numberOfResultsToReturn, wordSequenceSize)

	if filenames := flag.Args(); len(filenames) > 0 {
		slog.Debug(fmt.Sprintf("%d files passed to CLI", len(filenames)))

		for _, filename := range filenames {
			if err := cfg.AddExternalFile(filename); err != nil {
				panic(err)
			}
		}

		slog.Debug("Successfully added all files", "files", cfg.ExternalFilesList)
	} else {
		slog.Debug("No files passed to CLI, falling back to STDIN")

		scanner := bufio.NewScanner(os.Stdin)
		var userInput string
		for scanner.Scan() {
			userInput += scanner.Text()
		}
		if err := cfg.SetRawTextInput(userInput); err != nil {
			panic(err)
		}

		slog.Debug(fmt.Sprintf("Successfully added %d raw characters from STDIN", len(cfg.RawTextInput)))
	}

	// Run main application
	err := app.Run(cfg)
	if err != nil {
		slog.Error("Error encountered during main app execution!")
		panic(err)
	}
}
