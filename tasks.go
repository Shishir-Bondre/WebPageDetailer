package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func printHtmlVersion(doc *goquery.Document) {
	charset := doc.Find("charset").Contents().Text()
	fmt.Println("\n\n HTML Version: 5 ", charset)
}
