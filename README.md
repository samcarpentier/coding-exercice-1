# Trigrams Finder

Simple Go program that identifies the top-n most common _trigrams_ (sequences of 3 words) in a text.

## Requirements

* Built on MacOS Sonoma 14.4.1
* Git 2.41.0
* Go 1.22

```bash
brew install go
```

## Execution

### Local

* Install external dependencies:

```bash
go mod download
```

* Run the program:

```bash
# CLI usage:
# $ go run main.go -h
# Usage of trigrams:
#   -n int
#     	Number of results to return (default 100)
#   -s int
#     	The size of the word sequence to capture (default 3)
#   -v	Enables verbose logging

go run ./main.go -n 100 -s 3 ./texts/file1.txt ./texts/file2.txt
# or
cat ./texts/file1.txt ./texts/file2.txt | go run ./main.go -n 100 -s 3
# or
cat ./texts/* | go run ./main.go -n 100 -s 3
```

## Testing

### Local

* Run the tests

```bash
go test -v ./...
```

## Architecture & Design Overview

### High-Level Program Execution

* A configuration structure is instantiated from the CLI inputs provided by the user (see [`Config`](./config/config.go))
* If CLI was used with external file paths as positional arguments, files content is read from disk (see [`ExternalFilesReader`](./parser/external_files_reader.go))
* The raw text is then sanitized with the following rules (see [`TextSanitizer`](./parser/text_sanitizer.go)):
  * The whole text is converted to upper-case to make the count case insensitive
  * All punctuation signs are removed, leaving only alphanumerical characters, apostrophes (single quote) and spaces
  * Line breaks (`\n`) are converted to single spaces
  * Redundant spaces or tabs are removed
* The n-words sequences (n-grams) are indexed in a `map` structure and their occurrences throughout the sanitized text are counted (see [`NGramIndex`](./index/ngram_index.go))
* The map is converted to list of Key-Value pairs, sorted from the value (occurrences count)
* The final output is printed in a human-readable manner to the user

### How are Trigrams Identified

To ensure exhaustivity, n-words sequences are counted using a sliding window. Here is an example of sanitized text, with trigrams identified by square brackets

```
[WE'RE NO STRANGERS] TO LOVE
WE'RE [NO STRANGERS TO] LOVE
WE'RE NO [STRANGERS TO LOVE]
```

### Highlights

* Number of results returned by the program is configurable via CLI argument `-n`
* Size of the words sequences is configurable via CLI argument `-s`
* Unit tests have been written for all public functions of every package
* End-to-end tests have been written in the `app` package (see [`app_test.go`](./app/app_test.go))

## Next Steps


## Known Issues & Design Flaws

* The whole application is built with object-oriented (OO) paradigms even though Go is not an OO programming language per se. I chose to build the app this way since OO is the paradigm that I've used 95% of the time in the industry. However, I would be more than willing to get constructure criticism and tips on the best way to architect Go apps.

* At the moment, input files and STDIN input are read in their entirety upon program execution. This method of parsing doesn't scale to large files and could cause buffer overflows and/or abnormally high memory usage if in a production use-case.

* Hyphens on line-endings not handled; they are considered punctuation and stripped from the original text during the sanitization stage

* Handling of unicode characters, `\n`, `\r` and `\r\n` is simplistic. The bare minimum was done to make the program work with the dataset present in the [`texts/`](./texts/) folder.

* There is no STDIN timeout if no data is _piped_ into the process and no filenames are provided as positional arguments. The program will hang indefinitely if no input is provided.
