package controllers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: make region and language options
const (
	yahooURL = "https://yfapi.net/v6/finance/quote?symbols=%s&region=US&lang=en"
	api_key  = "4XKTWpU6YY2Y3N6zGKdip6iICRouIJmM83ePOUWD"
)

// Example request:
// https://yfapi.net/v6/finance/quote?region=US&lang=en&symbols=AAPL

// 4XKTWpU6YY2Y3N6zGKdip6iICRouIJmM83ePOUWD
// Yahoo finance API key
// TODO: make this an environment variable

func GetTicker(c *gin.Context) {
	symbol := c.Param("symbol")

	req, err := http.NewRequest("GET", fmt.Sprintf(yahooURL, symbol), nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	req.Header.Set("x-api-key", api_key)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	c.Data(http.StatusOK, gin.MIMEJSON, bodyBytes)
}
