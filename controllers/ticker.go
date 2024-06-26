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

func GetTicker(c *gin.Context) {
	if api_key == "" {
		initKey()
	}

	symbol := c.Param("symbol")

	// Create new HTTP request and add API key to the header
	req, err := initRequest("GET", fmt.Sprintf(yahooURL, symbol), nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	// Make HTTP request with the default client
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	defer response.Body.Close()

	// Read the body of the response as a slice of bytes and reset the io Reader
	bodyBytes, err := io.ReadAll(response.Body)
	response.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var badKey badResponse

	// Check if we got a bad response bc of no API key set
	d := json.NewDecoder(response.Body)
	err = d.Decode(&badKey)
	if err != nil || badKey.Message != "" {
		c.Data(http.StatusForbidden, gin.MIMEJSON, bodyBytes)
		return
	}

	c.Data(http.StatusOK, gin.MIMEJSON, bodyBytes)
}
