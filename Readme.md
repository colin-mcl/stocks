#  📈 Stocks


## 🌟 Highlights

- Simple command line and server API to view your favorite stock tickers!
- Gets up to date information from the Yahoo Finance API.
- Easy to download, set up, and use


## ℹ️ Overview
This is a personal interest project made by me (a graduating Tufts senior). I was interested in learning more about Golang, web services and APIs and I have a personal interest in finance so I created this app to combine those interests. The project is currently a work-in-progress but as of May 1, 2024 it provides a simple Go server programmed with the Gin package and net/http. It also provides an interactive command line interface called stocks_cli which can be used to make requests to the server's endpoint.

## 🚀 Usage
The project relies on the Yahoo Finance API provided by https://financeapi.net/. Make sure to have the enviroment variable STOCKS_API_KEY set to the API key obtained from that link. Additionally, set the STOCKS_URL environment variable to point to the url of the server, http://localhost:8080 by default.

```bash
$ go run server.go &
$ cd stocks_cli
$ go install
$ stocks_cli
                    STOCKS PROGRAM
Please enter 'get' followed by the stock ticker you would like to retrieve, or enter 'q' to quit
-------------------------------------------------------------------------
-> get TSLA
{
    "language": "en-US",
    "region": "US",
    "quoteType": "EQUITY",
    "typeDisp": "Equity",
    "quoteSourceName": "Nasdaq Real Time Price",
    "triggerable": true,
    "customPriceAlertConfidence": "HIGH",
    "currency": "USD",
    "marketState": "POST",
    "regularMarketChangePercent": -1.795064,
    "regularMarketPrice": 179.99,
    "exchange": "NMS",
    "shortName": "Tesla, Inc.",
    "longName": "Tesla, Inc.",
    "messageBoardId": "finmb_27444752",
    "exchangeTimezoneName": "America/New_York",
    "exchangeTimezoneShortName": "EDT",
    "gmtOffSetMilliseconds": -14400000,
    "market": "us_market",
    "esgPopulated": false,
    "hasPrePostMarketData": true,
    "firstTradeDateMilliseconds": 1277818200000,
    "priceHint": 2,
    "regularMarketChange": -3.2899933,
    "regularMarketTime": 1714593602,
    "regularMarketDayHigh": 185.86,
    "regularMarketDayRange": "179.01 - 185.86",
    "regularMarketDayLow": 179.01,
    "regularMarketVolume": 91800378,
    "regularMarketPreviousClose": 183.28,
    "bid": 179.96,
    "ask": 184.99,
    "bidSize": 2,
    "askSize": 4,
    "fullExchangeName": "NasdaqGS",
    "financialCurrency": "USD",
    "regularMarketOpen": 182,
    "averageDailyVolume3Month": 103347480,
    "averageDailyVolume10Day": 128319590,
    "fiftyTwoWeekLowChange": 41.190002,
    "fiftyTwoWeekLowChangePercent": 0.29675794,
    "fiftyTwoWeekRange": "138.8 - 299.29",
    "fiftyTwoWeekHighChange": -119.3,
    "fiftyTwoWeekHighChangePercent": -0.39861006,
    "fiftyTwoWeekLow": 138.8,
    "fiftyTwoWeekHigh": 299.29,
    "fiftyTwoWeekChangePercent": 14.114941,
    "earningsTimestamp": 1713907800,
    "earningsTimestampStart": 1721213940,
    "earningsTimestampEnd": 1721649600,
    "trailingAnnualDividendRate": 0,
    "trailingPE": 46.03325,
    "trailingAnnualDividendYield": 0,
    "epsTrailingTwelveMonths": 3.91,
    "epsForward": 3.33,
    "epsCurrentYear": 2.37,
    "priceEpsCurrentYear": 75.94515,
    "sharesOutstanding": 3189199872,
    "bookValue": 20.188,
    "fiftyDayAverage": 175.0682,
    "fiftyDayAverageChange": 4.9217987,
    "fiftyDayAverageChangePercent": 0.028113607,
    "twoHundredDayAverage": 221.1931,
    "twoHundredDayAverageChange": -41.203094,
    "twoHundredDayAverageChangePercent": -0.18627658,
    "marketCap": 574024122368,
    "forwardPE": 54.051052,
    "priceToBook": 8.915693,
    "sourceInterval": 15,
    "exchangeDataDelayedBy": 0,
    "averageAnalystRating": "2.8 - Hold",
    "tradeable": false,
    "cryptoTradeable": false,
    "displayName": "Tesla",
    "symbol": "TSLA"
}
```


## ⬇️ Installation

TODO: add installation instructions.

Requires Gin.