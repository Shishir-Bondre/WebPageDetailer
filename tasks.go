package main

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func printHtmlVersion(doc *goquery.Document) {
	charset := doc.Find("charset").Contents().Text()
	fmt.Println("\n\n HTML Version: 5 ", charset)
}

func printPageTitle(doc *goquery.Document) {
	pageTitle := doc.Find("title").Contents().Text()
	fmt.Println("\n PageTitle :", pageTitle)

}

func printTheNumberOfHeadings(doc *goquery.Document) {
	fmt.Printf("\n Heading count by level: ")

	// We have only h1-h6 headings in html
	for i := 1; i <= 6; i++ {
		tagName := "h1" + strconv.Itoa(i)
		fmt.Println("\n Tag Name: ", tagName, "Tag Count: ", getHeadingCount(doc, tagName))
	}
}

func getHeadingCount(doc *goquery.Document, heading string) int {
	headCount := 0
	doc.Find(heading).Each(func(i int, s *goquery.Selection) { headCount++ })
	return headCount
}
