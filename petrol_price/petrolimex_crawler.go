package petrol_price

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func CrawPetrolimex() {
	response, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	} else {
		log.Println(document)
	}
	document.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		// log.Println(s)
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
	defer response.Body.Close()
}
