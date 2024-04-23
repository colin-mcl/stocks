package main

import (
	"fmt"
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
	// reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\t\t\t STOCKS PROGRAM\n")
	fmt.Printf("\tPlease enter 'get' followed by the stock ticker you would like to retrieve\n")

}
