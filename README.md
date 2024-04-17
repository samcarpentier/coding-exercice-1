# Trigrams Finder

Simple Go program that identifies the top-n most common _trigrams_ (sequences of 3 words) in a text.

## Requirements

* MacOS Sonoma 14.4.1 (arm64)
* Go 1.22
* Git 2.41.0
* Podman 5.0.1

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
#
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

### Docker

A multi-stage [`Dockerfile`](./Dockerfile) has been built. The first stage uses a `golang` image to download dependencies and build a portable binary. The second stage is the runtime and uses a lightweight `alpine` Linux image. The binary is copied from the first stage onto the filesystem of the second stage and set as entrypoint.

* Build the image:

```bash
docker build -t trigrams:latest .
docker run trigrams:latest -h

# Usage of /usr/local/bin/trigrams:
#   -n int
#     	Number of results to return (default 100)
#   -s int
#     	The size of the word sequence to capture (default 3)
#   -v	Enables verbose logging
```

* Run the program inside a Docker container, with texts samples mounted via a volume:

```bash
docker run -v $PWD/texts:/home/appuser/texts:rw test:latest -n 5 -s 3 /home/appuser/texts/moby-dick.txt

# 2024/04/17 00:32:56 main.go:52:        INFO Using 1 specified files filenames=[/home/appuser/texts/moby-dick.txt]
# 2024/04/17 00:32:57 ngram_index.go:57: INFO Calculating the top-5 repeated 3 word sequences...
# 2024/04/17 00:32:57 app.go:30:         INFO Results compilation completed!

# Words Sequence       Count
# -------------------- ------
# THE SPERM WHALE      86
# OF THE WHALE         78
# THE WHITE WHALE      71
# ONE OF THE           64
# OUT OF THE           57
```

## Testing

### Local

* Run the tests:

```bash
go test -v ./...
```

### Dockerfile

* Run the tests:

```bash
docker run -ti -v $PWD:/data:rw golang:1.22.2-alpine3.19 sh
cd /data
go mod download
go test ./...

# /go # cd /data
# /data # go mod download
# /data # go test ./...
# ?   	trigrams	[no test files]
# ok  	trigrams/app	0.003s
# ok  	trigrams/config	0.002s
# ok  	trigrams/index	0.002s
# ok  	trigrams/parser	0.006s
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
* App can easily be packaged as a lightweight `alpine` Linux, non-root container

## Next Steps


## Known Issues & Design Flaws

* The whole application is built with object-oriented (OO) paradigms even though Go is not an OO programming language per se. I chose to build the app this way since OO is the paradigm that I've used 95% of the time in the industry. However, I would be more than willing to get constructure criticism and tips on the best way to architect Go apps.

* At the moment, input files and STDIN input are read in their entirety upon program execution. This method of parsing doesn't scale to large files and could cause buffer overflows and/or abnormally high memory usage if in a production use-case.

* Hyphens on line-endings not handled; they are considered punctuation and stripped from the original text during the sanitization stage

* Handling of unicode characters, `\n`, `\r` and `\r\n` is simplistic. The bare minimum was done to make the program work with the dataset present in the [`texts/`](./texts/) folder.

* There is no STDIN timeout if no data is _piped_ into the process and no filenames are provided as positional arguments. The program will hang indefinitely if no input is provided.

* Docker setup uses legacy multi-stage Dockerfile build instead of relying on the new `buildx` feature. This could be improved.

## If Given More Time...

1. Improve the way text is read from files or STDIN (line by line or group if lines by group of lines) instead of saving the whole raw text input in memory

2. Implement multi-threading to parallelize text reading and processing, with adequate locks on the input file and in-memory data structures to avoid concurrency issues when reading/writing data

3. Review the indexing and top-n results calculation methods in [`NGramIndex`](./index/ngram_index.go), including methods used from external libraries to verify time complexity of the operations being used. This could be an easy performance optimization of this code.

4. Make use of standard design patterns. Even if Object-oriented (OO) paradigms were used, no "proper" design patterns were implemented. The text sanitizer could be reimplemented as a pipe-and-filter pattern, the sorting algorithm could be reimplemented as a strategy pattern, etc.
