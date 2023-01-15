package fuel_price

import (
	"fmt"
	"sync"
	"time"
	"vn_oil_price/models"
)

var gasPriceCacheLock = &sync.Mutex{}

type DayGasPriceCache struct {
	priceMap map[string]models.GasolinePrice
}

type GasPriceCache struct {
	priceMap map[string]DayGasPriceCache
}

var gasPriceCacheInstance *GasPriceCache

func GetGasPriceCacheInstance() *GasPriceCache {
	if gasPriceCacheInstance == nil {
		gasPriceCacheLock.Lock()
		defer gasPriceCacheLock.Unlock()
		if gasPriceCacheInstance == nil {
			fmt.Println("Creating single instance for GasPriceCacheInstance")
			gasPriceCacheInstance = &GasPriceCache{
				priceMap: make(map[string]DayGasPriceCache),
			}
		}
	}

	return gasPriceCacheInstance
}

func (gasPriceCacheInstance *GasPriceCache) UpdateGasPriceByDate(date, name string, price models.GasolinePrice) {
	if dayPriceCache, ok := gasPriceCacheInstance.priceMap[date]; ok {
		dayPriceCache.priceMap[name] = price
	} else {
		newDayPriceCache := DayGasPriceCache{
			priceMap: make(map[string]models.GasolinePrice),
		}
		newDayPriceCache.priceMap[name] = price
		gasPriceCacheInstance.priceMap[date] = newDayPriceCache
	}
}

func (gasPriceCacheInstance *GasPriceCache) GetTodayGasPrice() (interface{}, error) {
	fmt.Println("--Start get today gas prices--")
	defer func() {
		fmt.Println("--Finish get today gas prices--")
	}()
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	currentTime := now.In(loc)
	today := currentTime.String()[0:10]

	if dayPriceCache, ok := gasPriceCacheInstance.priceMap[today]; ok {
		var r []models.GasolinePrice
		for _, element := range dayPriceCache.priceMap {
			r = append(r, element)
		}
		fmt.Println("Retrieve gas price from cache")
		return r, nil
	} else {
		err := fmt.Errorf("data is currently unavailable")
		return nil, err
	}
}
