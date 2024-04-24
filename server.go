package main

import (
	"github.com/colin-mcl/stocks/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/tickers/:symbol", controllers.GetTicker)

	// Runs the server on localhost:8080 by default
	router.Run()
}
