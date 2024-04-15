package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	TOP_N_RESULTS_TO_RETURN = 100
	NGRAM_SIZE              = 3
)

var ()

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if filenames := flag.Args(); len(filenames) > 0 {
		log.Printf("Found %d files: %v\n", len(filenames), filenames)
	} else {
		log.Println("Using stdin...")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}
