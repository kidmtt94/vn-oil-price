package petrol_price

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func CrawPetrolimex() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("div#vie_p5_PortletContent > div.list-table", func(e *colly.HTMLElement) {
		e.ForEach("div", func(i int, h *colly.HTMLElement) {
			// log.Println(h.Text)
		})

		divs := e.DOM.Children().Nodes
		for _, v := range divs {
			if v.Data == "section" {
				continue
			}
			child := v.FirstChild
			priceNameTag := child.FirstChild.FirstChild
			log.Println(priceNameTag.Data)
			areaOnePriceTag := child.NextSibling
			log.Println(areaOnePriceTag.FirstChild.Data)

			areaTwoPriceTag := areaOnePriceTag.NextSibling
			log.Println(areaTwoPriceTag.FirstChild.Data)

		}
		// e.Request.Post()
		// e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://public.petrolimex.com.vn/")
}
