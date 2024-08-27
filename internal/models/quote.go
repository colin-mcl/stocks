package models

// the quote type contains the most important fields from the Yahoo Finance
// API get quote endpoint, these can be changed to fit needs
type Quote struct {
	Symbol             string
	Region             string
	ShortName          string
	TimezoneShort      string
	Market             string
	Currency           string
	FiftyTwoWeekLow    float64
	FiftyTwoWeekHigh   float64
	RegularMarketPrice float64
}
