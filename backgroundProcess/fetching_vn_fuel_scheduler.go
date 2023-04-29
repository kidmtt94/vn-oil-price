package backgroundProcess

import (
	"time"
	"vn_oil_price/fuel_price"
)

func fetchingLatestPrices() {
	for {
		go fuel_price.FetchOilPrices()
		go fuel_price.FetchGasolinePrices()
		time.Sleep(120 * time.Minute)
	}
}
