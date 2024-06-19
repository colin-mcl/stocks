package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"syscall"

	"github.com/colin-mcl/stocks/client"
	"golang.org/x/term"
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

	user    *authenticatedUser
	running bool
}

// Represents an authenticated user on the stocks client
type authenticatedUser struct {
	email string

	accessToken string
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
		user:         nil,
	}
}

func (client *cmdClient) run() error {

	client.running = true
	fmt.Printf("\t\t\t\t\t  STOCKS PROGRAM\n")
	fmt.Printf("Please enter 'get' followed by the stock ticker you would like to retrieve, or enter 'q' to quit\n")
	fmt.Println("-----------------------------------------------------------------------------------------------")

	var err error = nil
	// infinite loop for user input
	for err == nil && client.running {
		err = client.loop()
	}

	return err
}

func (client *cmdClient) loop() error {

	if client.user != nil {
		fmt.Printf("%s ", client.user.email)
	}
	fmt.Print("-> ")

	// Get the next input line from stdin
	text, err := client.reader.ReadString('\n')

	// If error while getting line quit
	if err != nil {
		return err
	}

	// splits the input on whitespace
	words := strings.Fields(text)
	if len(words) == 0 || len(words) > 2 {
		client.errorLog.Printf("Invalid input: %s\n", text)
	}

	// as of 5/2/24 only option is get 'ticker' or q to quit
	switch strings.ToLower(words[0]) {
	case "quit":
		client.running = false
	case "get":
		err = client.stocksClient.GetQuote(words[1])
	case "create":
		err = client.createUser()
	case "login":
		err = client.login()
	case "user":
		err = client.getUser()
	case "logout":
		client.logout()
	default:
		client.errorLog.Println("Invalid command, see reference for command")
	}

	if err != nil {
		return err
	}
	return nil
}

func (client *cmdClient) getUser() error {
	if client.user == nil {
		client.errorLog.Printf("Error: must be logged in to get user info")
		return nil
	}

	user, err := client.stocksClient.GetUser(client.user.email, client.user.accessToken)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.Unauthenticated {
			client.errorLog.Print(err)
			return nil
		}
		return err
	}

	fmt.Println(user.String())
	return nil
}

func (client *cmdClient) logout() {
	if client.user != nil {
		client.user = nil
	} else {
		client.errorLog.Printf("Error: attempted to log out when no user is logged in")
	}
}

func (client *cmdClient) login() error {
	if client.user != nil {
		client.errorLog.Print("Error: already logged in, to log in again please log out first")
		return nil
	}

	fmt.Printf("Enter email:\n-> ")
	text, err := client.reader.ReadString('\n')
	if err != nil {
		return nil
	}

	words := strings.Fields(text)
	if len(words) != 1 {
		client.errorLog.Print("Error: too many arguments")
		return nil
	}

	email := words[0]

	fmt.Printf("Enter password:\n-> ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}

	token, err := client.stocksClient.LoginUser(email, string(bytePassword))

	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.Unauthenticated {
			client.errorLog.Print("Invalid credentials")
			return nil
		}
		return err
	}

	client.user = &authenticatedUser{email: email, accessToken: token}

	return nil
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
