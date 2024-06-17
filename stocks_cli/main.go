package main

/* reader.go

Colin Mclaughlin, April 2024

stocks project

reader.go provides the main method for the interactive stocks program,
the program can be run with go run reader.go.

It accepts the following options:
	- N/A for now


The following environment variables must be set:
	- STOCKS_URL=your_stocks_serverurl
		e.g. http://localhost:8080

Usage is as follows:
							STOCKS PROGRAM
	Please enter 'get' followed by the stock ticker you would like to retrieve

	-> get TSLA
	...
*/

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/colin-mcl/stocks/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var serverURL string

func main() {
	// gets the server URL to make requests or rpcs to
	serverURL = os.Getenv("STOCKS_URL")
	errorLog := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stderr, "[INFO] ", log.Ldate|log.Ltime)

	if serverURL == "" {
		errorLog.Fatal("Please set the STOCKS_URL environment variable" +
			" to your stock server address and restart the program.")
	}

	// Set up TLS credentials from ca-cert file
	creds, err := credentials.NewClientTLSFromFile("../cert/ca-cert.pem", "")
	if err != nil {
		errorLog.Fatal("failed to get certificate authority file: %w", err)
	}

	// Set up connection to the grpc server
	conn, err := grpc.NewClient(serverURL, grpc.WithTransportCredentials(creds))
	if err != nil {
		errorLog.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	cmdclient := newCmdClient(client.NewStocksClient(conn),
		errorLog, infoLog, bufio.NewReader(os.Stdin))

	err = cmdclient.run()
	if err != nil {
		errorLog.Fatal(err)
	}

}

// handleGetRequest
// helper function that makes the get ticker request to the server and unmarshals
// the json result into the result struct, returning a pointer to the struct
// and any errors that occured.
func handleHTTPGetRequest(ticker string) (*Result, error) {
	// Makes get request to HTTP endpoint set by environment variable
	url := fmt.Sprintf("%s/tickers/%s", serverURL, ticker)
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// deal with any error from bad status code
	if res.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(res.Body)
		return nil, errors.New(fmt.Sprintf("%s\n%s", res.Status, string(bodyBytes)))
	}

	// decode the json response into a result struct defined in structs.go
	var shell rspShell
	d := json.NewDecoder(res.Body)
	err = d.Decode(&shell)

	if err != nil {
		return nil, err
	} else if len(shell.QuoteResponse.Result) == 0 {
		return nil, errors.New(fmt.Sprintf("Ticker %s not found.\n", ticker))
	}

	return &shell.QuoteResponse.Result[0], nil
}
