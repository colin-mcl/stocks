package gapi

import (
	"io"
	"net/http"
	"os"
)

/* util.go

   Stocks Project

   Colin Mclaughlin, April 2024

   Util.go contains utility functions for the server controllers including
   getting the API key for Yahoo finance from the environment variables,
   checking for a bad response and more.

*/

var api_key string

func initKey() {
	// 4XKTWpU6YY2Y3N6zGKdip6iICRouIJmM83ePOUWD
	// Yahoo finance API key
	// Gets API key from environment variable
	if api_key == "" {
		api_key = os.Getenv("STOCKS_API_KEY")
	}
}

// init_request
// wrapper function for net/http.NewRequest that accepts the same parameters
// and adds the api key to the header
func initRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", api_key)
	return req, err
}
