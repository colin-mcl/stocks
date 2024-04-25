package main

import "encoding/json"

// structs.go defines all the structs needed to unmarshal the JSON format
// of stock ticker information that is returned by the server

// rspShell type contains a field called QuoteResponse of type quoteResponse
// with a JSON tag.
type rspShell struct {
	QuoteResponse quoteResponse `json:"quoteResponse"`
}

// quoteResponse contains the results of a ticker quote request:
// Result []Result is a slice of Result structs containing the response data
// Error is an interface{} type which can hold any type of value
type quoteResponse struct {
	Result []Result    `json:"result"`
	Error  interface{} `json"error"`
}

// Result contains the response data from a quote request of a specific ticker
// This struct was automatically generated using the website https://transform.tools/json-to-go
// TODO: add documentation for the fields
type Result struct {
	Language                          string  `json:"language"`
	Region                            string  `json:"region"`
	QuoteType                         string  `json:"quoteType"`
	TypeDisp                          string  `json:"typeDisp"`
	QuoteSourceName                   string  `json:"quoteSourceName"`
	Triggerable                       bool    `json:"triggerable"`
	CustomPriceAlertConfidence        string  `json:"customPriceAlertConfidence"`
	Currency                          string  `json:"currency"`
	MarketState                       string  `json:"marketState"`
	RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
	RegularMarketPrice                float64 `json:"regularMarketPrice"`
	Exchange                          string  `json:"exchange"`
	ShortName                         string  `json:"shortName"`
	LongName                          string  `json:"longName"`
	MessageBoardID                    string  `json:"messageBoardId"`
	ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
	ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
	GmtOffSetMilliseconds             int     `json:"gmtOffSetMilliseconds"`
	Market                            string  `json:"market"`
	EsgPopulated                      bool    `json:"esgPopulated"`
	HasPrePostMarketData              bool    `json:"hasPrePostMarketData"`
	FirstTradeDateMilliseconds        int64   `json:"firstTradeDateMilliseconds"`
	PriceHint                         int     `json:"priceHint"`
	RegularMarketChange               float64 `json:"regularMarketChange"`
	RegularMarketTime                 int     `json:"regularMarketTime"`
	RegularMarketDayHigh              float64 `json:"regularMarketDayHigh"`
	RegularMarketDayRange             string  `json:"regularMarketDayRange"`
	RegularMarketDayLow               float64 `json:"regularMarketDayLow"`
	RegularMarketVolume               int     `json:"regularMarketVolume"`
	RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose"`
	Bid                               float64 `json:"bid"`
	Ask                               float64 `json:"ask"`
	BidSize                           int     `json:"bidSize"`
	AskSize                           int     `json:"askSize"`
	FullExchangeName                  string  `json:"fullExchangeName"`
	FinancialCurrency                 string  `json:"financialCurrency"`
	RegularMarketOpen                 float64 `json:"regularMarketOpen"`
	AverageDailyVolume3Month          int     `json:"averageDailyVolume3Month"`
	AverageDailyVolume10Day           int     `json:"averageDailyVolume10Day"`
	FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange"`
	FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent"`
	FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange"`
	FiftyTwoWeekHighChange            float64 `json:"fiftyTwoWeekHighChange"`
	FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent"`
	FiftyTwoWeekLow                   float64 `json:"fiftyTwoWeekLow"`
	FiftyTwoWeekHigh                  float64 `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekChangePercent         float64 `json:"fiftyTwoWeekChangePercent"`
	EarningsTimestamp                 int     `json:"earningsTimestamp"`
	EarningsTimestampStart            int     `json:"earningsTimestampStart"`
	EarningsTimestampEnd              int     `json:"earningsTimestampEnd"`
	TrailingAnnualDividendRate        float64 `json:"trailingAnnualDividendRate"`
	TrailingPE                        float64 `json:"trailingPE"`
	TrailingAnnualDividendYield       float64 `json:"trailingAnnualDividendYield"`
	EpsTrailingTwelveMonths           float64 `json:"epsTrailingTwelveMonths"`
	EpsForward                        float64 `json:"epsForward"`
	EpsCurrentYear                    float64 `json:"epsCurrentYear"`
	PriceEpsCurrentYear               float64 `json:"priceEpsCurrentYear"`
	SharesOutstanding                 int64   `json:"sharesOutstanding"`
	BookValue                         float64 `json:"bookValue"`
	FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
	FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
	FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
	TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
	TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
	TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
	MarketCap                         int64   `json:"marketCap"`
	ForwardPE                         float64 `json:"forwardPE"`
	PriceToBook                       float64 `json:"priceToBook"`
	SourceInterval                    int     `json:"sourceInterval"`
	ExchangeDataDelayedBy             int     `json:"exchangeDataDelayedBy"`
	AverageAnalystRating              string  `json:"averageAnalystRating"`
	Tradeable                         bool    `json:"tradeable"`
	CryptoTradeable                   bool    `json:"cryptoTradeable"`
	DisplayName                       string  `json:"displayName"`
	Symbol                            string  `json:"symbol"`
}

func StructToString(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}

	return string(val), err
}
