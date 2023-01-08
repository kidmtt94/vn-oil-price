package server

import (
	"net/http"
	"vn_oil_price/petrol_price"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	// Ping test
	r.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "Hi")
	})

	r.GET("/prices/today", func(c *gin.Context) {
		result, err := petrol_price.GetPriceCacheInstance().GetTodayPrice()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"data": result})
		} else {
			c.Status(http.StatusNoContent)
		}
	})

	return r
}
