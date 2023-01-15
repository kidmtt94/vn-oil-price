package server

import (
	"net/http"
	"vn_oil_price/api"

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

	// Oil Price
	r.GET("/oil/prices/today", api.RetrieveTodayOilPrice)

	// Gasoline Price
	r.GET("/gasoline/prices/today", api.RetrieveTodayGasolinePrice)

	return r
}
