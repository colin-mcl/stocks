package controllers

/* ticker.go

Colin Mclaughlin, April 2024

stocks project

ticker.go is a part of the controller architecture of the stocks server.
This file handles requests to the /tickers/:symbol endpoint and returns
a JSON object with the results of the request to the Yahoo quotes endpoint.

The following environment variables must be set:
	- STOCKS_API_KEY=your_key
		This is the key to the https://financeapi.net/ free yahoo finance API

		TODO: add region and language options

*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// TODO: make region and language options
const (
	yahooURL = "https://yfapi.net/v6/finance/quote?symbols=%s&region=US&lang=en"
)

// badResponse indicates that the api key provided (or lack thereof) did not
// successfully work with the yahoo finance api and should be replaced
type badResponse struct {
	Message string `json:"message"`
	Hint    string `json:"hint"`
}

// Example request:
// https://yfapi.net/v6/finance/quote?region=US&lang=en&symbols=AAPL

var api_key string

// 4XKTWpU6YY2Y3N6zGKdip6iICRouIJmM83ePOUWD
// Yahoo finance API key
// TODO: make this an environment variable

func GetTicker(c *gin.Context) {
	if api_key == "" {
		api_key = os.Getenv("STOCKS_API_KEY")
	}

	symbol := c.Param("symbol")

	req, err := http.NewRequest("GET", fmt.Sprintf(yahooURL, symbol), nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	req.Header.Set("x-api-key", api_key)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	response.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var badKey badResponse
	d := json.NewDecoder(response.Body)
	err = d.Decode(&badKey)
	if err != nil || badKey.Message != "" {
		c.Data(http.StatusForbidden, gin.MIMEJSON, bodyBytes)
		return
	}

	c.Data(http.StatusOK, gin.MIMEJSON, bodyBytes)
}
