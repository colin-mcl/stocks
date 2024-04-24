package main

import (
	"github.com/colin-mcl/stocks/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/tickers", controllers.GetTicker)

}
