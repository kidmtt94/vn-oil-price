package main

import (
	"vn_oil_price/petrol_price"
)

func main() {
	// r := server.SetupRouter()
	petrol_price.CrawPetrolimex()
	// Listen and Server in 0.0.0.0:8080
	// r.Run(":8081")

}
