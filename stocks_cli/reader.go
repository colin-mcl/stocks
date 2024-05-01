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
	"strings"
)

var serverURL string

func main() {
	reader := bufio.NewReader(os.Stdin)
	serverURL = os.Getenv("STOCKS_URL")

	if serverURL == "" {
		fmt.Fprintln(os.Stderr, "Please set the STOCKS_URL environment variable"+
			" to your stock server address and restart the program.")
		os.Exit(1)
	}

	fmt.Printf("\t\t\t\t\t  STOCKS PROGRAM\n")
	fmt.Printf("Please enter 'get' followed by the stock ticker you would like to retrieve, or enter 'q' to quit\n")
	fmt.Println("-----------------------------------------------------------------------------------------------")

	for {
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

		if len(words) == 1 {
			if strings.ToLower(words[0]) == "q" {
				break
			} else {
				fmt.Fprintln(os.Stderr, "Not enough arguments, please provide ticker name.")
				continue
			}
		}

		res, err := handleGetRequest(strings.ToUpper(words[1]))

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}

		s, err := StructToString(res)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}

		fmt.Printf("%s\n", s)
	}

}

func handleGetRequest(ticker string) (*Result, error) {
	url := fmt.Sprintf("%s/tickers/%s", serverURL, ticker)
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(res.Body)
		return nil, errors.New(fmt.Sprintf("%s\n%s", res.Status, string(bodyBytes)))
	}

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
