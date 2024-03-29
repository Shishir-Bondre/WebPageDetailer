package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	text := getInputFromTerminal()
	doc := getHTMLPageFromInternet(text)
	printHtmlVersion(doc)
	printPageTitle(doc)
	printTheNumberOfHeadings(doc)
	printInternalAndExternalLinks(doc)
	printInacessibleLinks(doc)
	printTrueIfLoginFormExists(doc)
	fmt.Println(doc)

}

func getInputFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter webiste url for details : \n e.g. https://www.home24.com/  \n")
	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
		return ""
	}

	text = strings.Replace(text, "\n", "", -1)
	return text
}

func getHTMLPageFromInternet(text string) *goquery.Document {
	doc, err := goquery.NewDocument(text)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
