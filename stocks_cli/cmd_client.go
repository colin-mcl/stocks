package main

/* cmd_client.go

Contains implementation for the private cmdClient struct which wraps all
available commands and implementation details for the command line interface.

Supported commands:
	To see the list of support commands
	    -> help

	To get a quote for any stock symbol (e.g. TSLA, AAPL) and print in formatted
	form
		-> get 'symbol'

	To create a user account which can be used to login and see portfolios
   		-> create

	To login to a created user account and see your portfolio
		-> login

	To get information on the current, LOGGED IN, user
		-> user

	To logout of the current user account
		-> logout

	To quit the program
		-> quit

	TODO: add these to README
	More commands to come...
*/

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"syscall"

	client "github.com/colin-mcl/stocks/pkg/v1/stocks_client"
	"golang.org/x/term"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// cmdClient struct wraps the stocksClient, loggers, user and reader so that
// simple functions can be called to perform the implemented commands
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

// Runs the command client
func (client *cmdClient) run() error {

	client.running = true
	fmt.Printf("\t\t\t\t\t  STOCKS PROGRAM\n")
	fmt.Printf("Run `help` to see available commands\n")
	fmt.Println("-----------------------------------------------------------------------------------------------")

	var err error = nil
	// infinite loop for reading user commands
	for err == nil && client.running {
		err = client.loop()
	}

	return err
}

func (client *cmdClient) help() {
	fmt.Println(`
***** COMMANDS ******
To see the list of support commands
  -> help

To get a quote for any stock symbol (e.g. TSLA, AAPL) and print in formatted
form
  -> get 'symbol'

To create a user account which can be used to login and see portfolios
  -> create

To login to a created user account and see your portfolio
  -> login

To get information on the current, LOGGED IN, user
  -> user

To logout of the current user account
  -> logout

To quit the program
  -> quit`)
}

// Loops infinitely until a fatal error occurs or the user quits
func (client *cmdClient) loop() error {

	// print current logged in user before prompt
	if client.user != nil {
		fmt.Printf("%s ", client.user.email)
	}
	fmt.Print("-> ")

	text, err := client.reader.ReadString('\n')
	if err != nil {
		return err
	}

	// splits the input on whitespace
	words := strings.Fields(text)
	if len(words) == 0 {
		return nil
	}
	command := strings.ToLower(words[0])

	switch command {
	case "help":
		client.help()
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

// Gets information of the current logged in user, or prints an error if no
// user is currently logged in
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

// Logs out the current user or prints an error message if no user is logged in
func (client *cmdClient) logout() {
	if client.user != nil {
		client.user = nil
	} else {
		client.errorLog.Printf("Error: attempted to log out when no user is logged in")
	}
}

// Logs into a user account by prompting for email and password and making
// request to the stocksClient
func (client *cmdClient) login() error {
	if client.user != nil {
		client.errorLog.Print("Error: already logged in, to log in again please log out first")
		return nil
	}

	fmt.Printf("Enter email: ")
	text, err := client.reader.ReadString('\n')
	if err != nil {
		return nil
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		client.errorLog.Print("no email entered")
		return nil
	}
	email := words[0]

	// Get password from stdin without displaying the text
	fmt.Printf("Enter password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}

	// authenticate user
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
	fmt.Println()
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
