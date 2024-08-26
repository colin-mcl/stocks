package main

/* main.go

Colin Mclaughlin, April 2024

stocks project

main launches the interactive command line interface of the stocks program by
connecting to the grpc server and intializing and running a command client
struct.

It accepts the following options:
	- N/A for now


The following configurations must be in place:
	- env variable STOCKS_URL=your_stocks_serverurl
		e.g. http://localhost:9090
	- TLS trusted certificates file (ca-cert.pem) must be stored in a folder
	   named cert one level above this folder
	TODO: change this to environment variable

See cmd_client.go for documentation on commands which are accepted
*/

import (
	"bufio"
	"log"
	"os"

	"github.com/colin-mcl/stocks/internal/cli"
	"github.com/colin-mcl/stocks/pkg/v1/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// var serverURL string

func main() {
	// Create logger objects on stderr
	errorLog := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stderr, "[INFO] ", log.Ldate|log.Ltime)

	// gets the server URL to make requests or rpcs to
	serverURL := os.Getenv("STOCKS_URL")

	if serverURL == "" {
		errorLog.Fatal("Please set the STOCKS_URL environment variable" +
			" to your stock server address and restart the program.")
	}

	certPath := os.Getenv("CERT_PATH")

	// Set up TLS credentials from ca-cert file
	creds, err := credentials.NewClientTLSFromFile(certPath, "")
	if err != nil {
		errorLog.Fatal("failed to get certificate authority file: %w", err)
	}

	// Create connection to GRPC server with TLS credentials
	conn, err := grpc.NewClient(serverURL, grpc.WithTransportCredentials(creds))
	if err != nil {
		errorLog.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	CLI := cli.NewCLI(client.NewStocksClient(conn),
		errorLog, infoLog, bufio.NewReader(os.Stdin))

	// Run command client and log any fatal errors
	err = CLI.Run()
	if err != nil {
		errorLog.Fatal(err)
	}

}

// NOT CURRENTLY IN USE:
//
// handleGetRequest
// helper function that makes the get ticker request to the server and unmarshals
// the json result into the result struct, returning a pointer to the struct
// and any errors that occured.
// func handleHTTPGetRequest(ticker string) (*Result, error) {
// 	// Makes get request to HTTP endpoint set by environment variable
// 	url := fmt.Sprintf("%s/tickers/%s", serverURL, ticker)
// 	res, err := http.Get(url)

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	// deal with any error from bad status code
// 	if res.StatusCode != http.StatusOK {
// 		bodyBytes, _ := io.ReadAll(res.Body)
// 		return nil, errors.New(fmt.Sprintf("%s\n%s", res.Status, string(bodyBytes)))
// 	}

// 	// decode the json response into a result struct defined in structs.go
// 	var shell rspShell
// 	d := json.NewDecoder(res.Body)
// 	err = d.Decode(&shell)

// 	if err != nil {
// 		return nil, err
// 	} else if len(shell.QuoteResponse.Result) == 0 {
// 		return nil, errors.New(fmt.Sprintf("Ticker %s not found.\n", ticker))
// 	}

// 	return &shell.QuoteResponse.Result[0], nil
// }
