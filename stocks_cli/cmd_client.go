package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/colin-mcl/stocks/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// cmd_client contains a command client struct which wraps all of the
// required fields to run an interactive command line interface with the
// stocks client

type cmdClient struct {
	// stocksClient for interfacing with the stocks service
	stocksClient *client.StocksClient

	errorLog *log.Logger
	infoLog  *log.Logger

	// stdin reader
	reader *bufio.Reader
}

func newCmdClient(
	c *client.StocksClient,
	errorLog *log.Logger,
	infoLog *log.Logger,
	r *bufio.Reader) *cmdClient {
	return &cmdClient{
		stocksClient: c,
		errorLog:     errorLog,
		infoLog:      infoLog,
		reader:       r,
	}
}

func (client *cmdClient) run() error {

	fmt.Printf("\t\t\t\t\t  STOCKS PROGRAM\n")
	fmt.Printf("Please enter 'get' followed by the stock ticker you would like to retrieve, or enter 'q' to quit\n")
	fmt.Println("-----------------------------------------------------------------------------------------------")

	running := true
	var err error
	// infinite loop for user input
	for running {
		running, err = client.loop()

		if err != nil {
			return err
		}
	}

	return nil
}

func (client *cmdClient) loop() (bool, error) {

	fmt.Print("-> ")

	// Get the next input line from stdin
	text, err := client.reader.ReadString('\n')

	// If error while getting line quit
	if err != nil {
		return false, err
	}

	// splits the input on whitespace
	words := strings.Fields(text)
	if len(words) == 0 || len(words) > 2 {
		client.errorLog.Printf("Invalid input: %s\n", text)
	}

	// as of 5/2/24 only option is get 'ticker' or q to quit
	switch strings.ToLower(words[0]) {
	case "quit":
		return false, nil
	case "get":
		err := client.stocksClient.GetQuote(words[1])
		if err != nil {
			return false, err
		}
	case "create":
		err := client.createUser()
		if err != nil {
			return false, err
		}
	default:
		client.errorLog.Println("Invalid command, see reference for command")
	}

	// if len(words) == 1 {
	// 	if strings.ToLower(words[0]) == "q" ||
	// 		strings.ToLower(words[0]) == "quit" {
	// 		return false, nil
	// 	} else {
	// 		errorLog.Println("Not enough arguments, please provide ticker name.")
	// 		return true, nil
	// 	}
	// }

	// client.GetQuote(words[1])
	return true, nil
}

// createUser gets the appropriate fields from reader to create a user on
// the stocks client
func (client *cmdClient) createUser() error {
	fmt.Printf("Enter the following fields separated by spaces: firstname, lastname, username, email, password:\n-> ")
	text, err := client.reader.ReadString('\n')
	if err != nil {
		return err
	}

	words := strings.Fields(text)
	if len(words) != 5 {
		client.errorLog.Printf("Error: incorrect number of fields for create user")
		return nil
	}

	id, err := client.stocksClient.CreateUser(words[0], words[1], words[2], words[3], words[4])

	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			client.errorLog.Print("user already exists")
			return nil
		} else {
			return err
		}
	}

	fmt.Printf("user created with id: %d\n", id)
	return nil
}
