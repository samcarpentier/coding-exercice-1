# Trigrams Finder

Simple Go program that identifies the top-n most common _trigrams_ (sequences of 3 words) in a text.

## Requirements

* Go 1.22

```bash
brew install go
```

* Some samples of long text to parse with the program, under the `text/` folder

```bash
curl https://raw.githubusercontent.com/Axosoft/moby-git/main/full-text-moby-dick.txt -o ./texts/moby-dick.txt
curl https://raw.githubusercontent.com/kuemit/txt_book/master/examples/alice_in_wonderland.txt -o ./texts/alice-in-wonderland.txt
```

## How to Run

* To install external dependencies:

```bash
go mod download
```

* To run the program:

```bash
go run ./main.go -- ./texts/file1.txt ./texts/file2.txt
# or
cat ./texts/file1.txt | go run ./main.go
```

* To run all tests:

```bash
go test -v ./...
```

## Next Steps


## Known Issues & Design Flaws

* The whole application is built with object-oriented paradigms even though Go is not an OO programming language perse. I chose to build the app this way since OO is the paradigm that I've used 95% of the time in the industry. However, I would be more than willing to get constructure criticism and tips on the best way to architect Go apps.

* At the moment, input files or STDIN input are read in their entirety upon program execution. This method of parsing doesn't scale to large files and could cause buffer overflows and/or abnormally high memory usage if in a production use-case.

* Hyphens on line-endings not handled
