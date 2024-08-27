package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/colin-mcl/stocks/internal/models"
)

// TODO: make region and language options
// 4XKTWpU6YY2Y3N6zGKdip6iICRouIJmM83ePOUWD - API key
const yahooURL = "https://yfapi.net/v6/finance/quote?symbols=%s&region=US&lang=en"

// GetQuote
//
// Gets the quote indicated in the stock ticker symbol from the yahoo finance
// api and returns the result
func (uc *UseCase) GetQuote(symbol string) (*models.Quote, error) {
	// create http request object with yahooo url and api key
	req, err := http.NewRequest("GET",
		fmt.Sprintf(yahooURL, url.QueryEscape(symbol)), nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", os.Getenv("STOCKS_API_KEY"))

	// execute http request using http.DefaultClient
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	// read the http response body contents as byte slice
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// unmarshal body contents into json map in data
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	// check if we failed to get a response due to bad api key
	if data["message"] != nil {
		return nil, ErrBadKey
	}

	// json response structure example:
	/*
		{
		 "quoteResponse: {
			    "result: [
				  {
				    "symbol": "TSLA"
					}
				]
				}
			}
	*/
	qR, ok := data["quoteResponse"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to unmarshal quote response")
	}

	result, ok := qR["result"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to unmarshal result")
	}
	if len(result) == 0 {
		return nil, ErrBadSymbol
	}

	quote, ok := result[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to unmarshal quote")
	}

	q := &models.Quote{
		Symbol:             quote["symbol"].(string),
		Region:             quote["region"].(string),
		ShortName:          quote["shortName"].(string),
		TimezoneShort:      quote["exchangeTimezoneShortName"].(string),
		Market:             quote["market"].(string),
		Currency:           quote["currency"].(string),
		FiftyTwoWeekLow:    quote["fiftyTwoWeekLow"].(float64),
		FiftyTwoWeekHigh:   quote["fiftyTwoWeekHigh"].(float64),
		RegularMarketPrice: quote["regularMarketPrice"].(float64),
	}
	return q, nil
}
