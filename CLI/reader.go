package main

import (
	"bufio"
	"fmt"
	"log"
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

		if len(words) == 1 && strings.ToLower(words[0]) == "q" {
			break
		}
	}

}
