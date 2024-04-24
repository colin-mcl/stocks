package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	yahooURL = "https://query1.finance.yahoo.com/v8/finance/chart/%s"
)

func GetTicker(c *gin.Context) {
	symbol := c.Param("symbol")
	c.String(http.StatusOK, "%s\n", symbol)
}
