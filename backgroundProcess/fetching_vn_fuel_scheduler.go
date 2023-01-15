package backgroundProcess

import (
	"time"
	"vn_oil_price/fuel_price"
)

func fetchingLatestPrices() {
	for {
		fuel_price.FetchOilPrices()
		fuel_price.FetchGasolinePrices()
		time.Sleep(120 * time.Minute)
	}
}
