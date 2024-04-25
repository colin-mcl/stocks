package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

/* reader.go

Colin Mclaughlin, April 2024

stocks project

reader.go provides the main method for the interactive stocks program,
the program can be run with go run reader.go.

It accepts the following options:
	- N/A for now

Usage is as follows:
							STOCKS PROGRAM
	Please enter 'get' followed by the stock ticker you would like to retrieve

	-> get TSLA
	...
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

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
				fmt.Fprint(os.Stderr, "Not enough arguments, please provide ticker name.\n")
				continue
			}
		}

		err = handleRequest(strings.ToUpper(words[1]))

		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
		}
	}

}

// TODO: make this an env variable?
func handleRequest(ticker string) error {
	url := fmt.Sprintf("http://localhost:8080/tickers/%s", ticker)
	res, err := http.Get(url)

	if err != nil {
		return err
	}

	var shell rspShell
	d := json.NewDecoder(res.Body)
	err = d.Decode(&shell)
	if err != nil {
		return err
	} else if len(shell.QuoteResponse.Result) == 0 {
		return errors.New(fmt.Sprintf("Ticker %s not found.\n", ticker))
	}

	s, err := StructToString(&shell.QuoteResponse.Result[0])

	if err != nil {
		return err
	}

	fmt.Printf("%s\n", s)
	return nil
}
