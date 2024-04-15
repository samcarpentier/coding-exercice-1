# Trigrams Finder

Simple Golang program that identifies the top-n most common _trigrams_ (sequences of 3 words) in a text.

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
go test ./...
```

## Next Steps


## Known Issues

