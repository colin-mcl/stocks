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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/colin-mcl/stocks/pb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverURL string

func main() {
	// gets the server URL to make requests or rpcs to
	serverURL = os.Getenv("STOCKS_URL")

	if serverURL == "" {
		fmt.Fprintln(os.Stderr, "Please set the STOCKS_URL environment variable"+
			" to your stock server address and restart the program.")
		os.Exit(1)
	}

	// Set up connection to the grpc server
	conn, err := grpc.NewClient(serverURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewStocksClient(conn)

	fmt.Printf("\t\t\t\t\t  STOCKS PROGRAM\n")
	fmt.Printf("Please enter 'get' followed by the stock ticker you would like to retrieve, or enter 'q' to quit\n")
	fmt.Println("-----------------------------------------------------------------------------------------------")

	// infinite loop for user input
	for loop(c) {
	}

}

func loop(c pb.StocksClient) bool {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")

	// Get the next input line from stdin
	text, err := reader.ReadString('\n')

	// If error while getting line quit
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// splits the input on whitespace
	words := strings.Fields(text)
	if len(words) == 0 || len(words) > 2 {
		fmt.Fprintf(os.Stderr, "Invalid input: %s\n", text)
	}

	// as of 5/2/24 only option is get 'ticker' or q to quit
	if len(words) == 1 {
		if strings.ToLower(words[0]) == "q" {
			return false
		} else {
			fmt.Fprintln(os.Stderr, "Not enough arguments, please provide ticker name.")
			return true
		}
	}

	res, err := c.GetQuote(context.Background(), &pb.GetQuoteRequest{Symbol: strings.ToUpper(words[1])})

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get ticker: %v\n", err)
		return true
	}
	fmt.Println(proto.MarshalTextString(res))

	return true
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
