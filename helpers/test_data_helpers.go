package helpers

import (
	"bufio"
	"log"
	"os"

	"github.com/SethHafferkamp/GoRSS/feedparsing"
)

func RegisterTestData() {
	readFile, err := os.Open("test_feeds.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	for _, eachline := range fileTextLines {
		feedparsing.RegisterFeed(eachline)
		// fmt.Println(eachline)
	}
}
