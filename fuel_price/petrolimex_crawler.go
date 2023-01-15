package fuel_price

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
	"vn_oil_price/models"

	"github.com/gocolly/colly"
)

func FetchOilPrices() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("div#vie_p5_PortletContent > div.list-table", func(e *colly.HTMLElement) {
		now := time.Now()
		loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
		currentTime := now.In(loc)
		today := currentTime.String()[0:10]
		log.Println("Start crawling oil prices", currentTime)
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

func FetchGasolinePrices() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("div#vie_p8_PortletContent > div.list-table", func(e *colly.HTMLElement) {
		now := time.Now()
		loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
		currentTime := now.In(loc)
		today := currentTime.String()[0:10]
		log.Println("Start crawling gasoline prices", currentTime)
		divs := e.DOM.Children().Nodes
		for _, v := range divs {
			if v.Data == "section" {
				continue
			}
			child := v.FirstChild
			locationNameTag := child.FirstChild.FirstChild
			locationValue := locationNameTag.Data
			cyclinder12PriceTag := child.NextSibling
			cylinder12Value := cyclinder12PriceTag.FirstChild.Data
			cyclinder48Tag := cyclinder12PriceTag.NextSibling
			cylinder48Value := cyclinder48Tag.FirstChild.Data

			info := models.GasolinePrice{
				Location:     locationValue,
				Cylinder12KG: cylinder12Value,
				Cylinder48KG: cylinder48Value,
			}
			GetGasPriceCacheInstance().UpdateGasPriceByDate(today, locationValue, info)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	})
	c.Visit("https://www.pgas.petrolimex.com.vn/")
}
