package petrol_price

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

type petrolPriceCache struct {
	priceMap map[string]dayPriceCache
}

type dayPriceCache struct {
	priceMap map[string]string
}

var singleInstance *single

func GetInstance() *single {
	a := models.
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
