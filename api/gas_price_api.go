package api

import (
	"net/http"
	"vn_oil_price/fuel_price"

	"github.com/gin-gonic/gin"
)

// RetrieveTodayGasolinePrice return data gas prices today in json format
func RetrieveTodayGasolinePrice(c *gin.Context) {
	result, err := fuel_price.GetGasPriceCacheInstance().GetTodayGasPrice()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"data": result})
	} else {
		c.Status(http.StatusNoContent)
	}
}
