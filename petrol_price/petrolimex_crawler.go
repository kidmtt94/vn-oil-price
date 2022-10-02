package petrol_price

import (
	"log"
	"time"
	"vn_oil_price/models"

	"github.com/gocolly/colly"
)

func CrawPetrolimex() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("div#vie_p5_PortletContent > div.list-table", func(e *colly.HTMLElement) {
		now := time.Now()
		loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
		currentTime := now.In(loc)
		today := currentTime.String()[0:10]
		log.Println("Start crawling", currentTime)
		divs := e.DOM.Children().Nodes
		for _, v := range divs {
			if v.Data == "section" {
				continue
			}
			child := v.FirstChild
			priceNameTag := child.FirstChild.FirstChild
			oilName := priceNameTag.Data
			areaOnePriceTag := child.NextSibling
			areaOneValue := areaOnePriceTag.FirstChild.Data
			areaTwoPriceTag := areaOnePriceTag.NextSibling
			areaTwoValue := areaTwoPriceTag.FirstChild.Data

			info := models.OilPrice{
				Name:         oilName,
				AreaOnePrice: areaOneValue,
				AreaTwoPrice: areaTwoValue,
			}
			GetPriceCacheInstance().UpdatePriceByDate(today, oilName, info)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.Visit("https://public.petrolimex.com.vn/")
}
