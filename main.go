package main

import (
	"fmt"
	"github.com/headzoo/surf"
	"strconv"
	"strings"
	"time"
)

func main() {

	for pageNumber := 1; pageNumber <= 3434; pageNumber++ { // iterate thru pages 1-3434
		for elementNumber := 1; elementNumber <= 20; elementNumber++ { // iterate thru the 20 elements from each page
			fmt.Println(getListing("https://obchody.heureka.cz/?f="+strconv.Itoa(pageNumber), elementNumber)) // output the scraped data
		}
	}

}

func getListing(url string, listingNumber int) (string, string, string) {
	bow := surf.NewBrowser()
	err := bow.Open(url)
	if err != nil {
		panic(err)
	}

	bow.SetTimeout(10 * time.Second)

	var urlListing, title, urlWebsite string

	bow.Click(".c-shops-table__row:nth-child(" + strconv.Itoa(listingNumber) + ") .e-button--simple") // grab listing URL
	urlListing = bow.Url().String()
	title = strings.TrimRight(bow.Title(), " — Heureka.cz") // grab company name
	bow.Back()

	bow.Click(".c-shops-table__row:nth-child(" + strconv.Itoa(listingNumber) + ") .e-button:nth-child(2)") // grab website URL
	urlWebsite = bow.Url().String()

	if strings.Contains(title, "Internetové obchody") {
		title = "" // timed out or some error occured
	}

	if strings.Contains(urlListing, "https://obchody.heureka.cz/?f=") {
		urlListing = "" // timed out or some error occured
	}

	// if website doesn't load within 10seconds the page url will still be on listing url of heureka.cz and not the actual website
	if strings.Contains(urlWebsite, "heureka.cz") {
		urlWebsite = "" // timed out or some error occured
	}

	return title, urlListing, urlWebsite
}
