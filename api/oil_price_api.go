package api

import (
	"net/http"
	"vn_oil_price/fuel_price"

	"github.com/gin-gonic/gin"
)

// RetrieveTodayOilPrice return data oil prices today in json format
func RetrieveTodayOilPrice(c *gin.Context) {
	result, err := fuel_price.GetPriceCacheInstance().GetTodayPrice()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"data": result})
	} else {
		c.Status(http.StatusNoContent)
	}
}
