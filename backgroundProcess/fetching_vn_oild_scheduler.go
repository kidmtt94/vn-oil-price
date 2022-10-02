package backgroundProcess

import (
	"time"
	"vn_oil_price/petrol_price"
)

func fetchingPetrol() {
	for {
		petrol_price.CrawPetrolimex()
		time.Sleep(120 * time.Minute)
	}
}
