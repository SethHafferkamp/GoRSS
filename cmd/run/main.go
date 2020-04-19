package main

import (
	"fmt"
	"hello/feedparsing"
)

func main() {
	fmt.Println("Scraping and parsing all feeds")
	feedparsing.ParseAllFeeds(32)
	fmt.Println("All Done!!")
}
