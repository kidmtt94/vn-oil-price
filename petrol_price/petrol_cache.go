package petrol_price

import (
	"fmt"
	"sync"
	"time"
	"vn_oil_price/models"
)

var lock = &sync.Mutex{}

type DayPriceCache struct {
	priceMap map[string]models.OilPrice
}

type PetrolPriceCache struct {
	priceMap map[string]DayPriceCache
}

var priceCacheInstance *PetrolPriceCache

func GetPriceCacheInstance() *PetrolPriceCache {
	if priceCacheInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if priceCacheInstance == nil {
			fmt.Println("Creating single instance now.")
			priceCacheInstance = &PetrolPriceCache{
				priceMap: make(map[string]DayPriceCache),
			}
		}
	}

	return priceCacheInstance
}

func (priceCacheInstance *PetrolPriceCache) UpdatePriceByDate(date, name string, price models.OilPrice) {
	if dayPriceCache, ok := priceCacheInstance.priceMap[date]; ok {
		dayPriceCache.priceMap[name] = price
	} else {
		newDayPriceCache := DayPriceCache{
			priceMap: make(map[string]models.OilPrice),
		}
		newDayPriceCache.priceMap[name] = price
		priceCacheInstance.priceMap[date] = newDayPriceCache
	}
}

func (priceCacheInstance *PetrolPriceCache) GetTodayPrice() (interface{}, error) {
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	currentTime := now.In(loc)
	today := currentTime.String()[0:10]

	if dayPriceCache, ok := priceCacheInstance.priceMap[today]; ok {
		return dayPriceCache.priceMap, nil
	} else {
		err := fmt.Errorf("data is currently unavailable")
		return nil, err
	}
}
