package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

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

func printInternalAndExternalLinks(doc *goquery.Document) {
	externalLinks := getExternalLinksSize(doc)
	internalLinks := getInternalLinksSize(doc)
	fmt.Println("\n Amount of External Links: ", externalLinks)
	fmt.Println("\n Amount of Internal Links: ", internalLinks)
}

func getExternalLinksSize(doc *goquery.Document) int {
	externalLinks := getExternalLinks(doc)
	return len(externalLinks)
}

func getExternalLinks(doc *goquery.Document) []string {
	var external = []string{}
	doc.Find("body a").Each(func(index int, element *goquery.Selection) {
		link, _ := element.Attr("href")
		if strings.Contains(link, "http") {
			external = append(external, link)
		}
	})
	return external
}

func getInternalLinksSize(doc *goquery.Document) int {
	internalLinksCounter := 0
	doc.Find("body a").Each(func(index int, element *goquery.Selection) {
		link, _ := element.Attr("href")
		if !strings.Contains(link, "http") {
			internalLinksCounter++
		}
	})
	return internalLinksCounter
}

func printInacessibleLinks(doc *goquery.Document) {
	inacessible := urlCallCount(doc)
	fmt.Println("\n Checking links..... ")
	fmt.Println(" Amount of inacessible links: ", inacessible)

}

func urlCallCount(doc *goquery.Document) int {

	links := getExternalLinks(doc)
	count := 0
	var wg sync.WaitGroup

	// Parallel processing, can be delegated to the different function call
	// TODO: Add function for parallel async processing

	for _, link := range links {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			_, err := http.Get(url)
			if err != nil {
				count++
				fmt.Println("    Invalid link : ", url)
			}
		}(link)
	}
	wg.Wait()

	return count
}
