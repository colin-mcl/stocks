package cli

/* cli.go

Contains implementation for the private CLI struct which wraps all
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

	"github.com/colin-mcl/stocks/pkg/v1/client"
	"golang.org/x/term"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CLI struct wraps the stocksClient, loggers, user and reader so that
// simple functions can be called to perform the implemented commands
type CLI struct {
	// stocksClient for interfacing with the stocks service
	stocksClient *client.StocksClient

	errorLog *log.Logger
	infoLog  *log.Logger

	// stdin reader
	reader *bufio.Reader

	user    *authenticatedUser
	running bool
}

// Represents an authenticated user on the stocks cli
type authenticatedUser struct {
	email string

	accessToken string
}

// NewCLI
//
// Given the parameters needed to run the CLI, create a new CLI instance
// and return its pointer
func NewCLI(
	c *client.StocksClient,
	errorLog *log.Logger,
	infoLog *log.Logger,
	r *bufio.Reader) *CLI {
	return &CLI{
		stocksClient: c,
		errorLog:     errorLog,
		infoLog:      infoLog,
		reader:       r,
		user:         nil,
	}
}

// Run
//
// Runs the CLI
func (c *CLI) Run() error {

	c.running = true
	fmt.Printf("\t\t\t\t\t  STOCKS PROGRAM\n")
	fmt.Printf("Run `help` to see available commands\n")
	fmt.Println("-----------------------------------------------------------------------------------------------")

	var err error
	// infinite loop for reading user commands
	for c.running {
		err = c.loop()
		if err != nil {
			c.errorLog.Println(err)
		}
	}

	return err
}

// help
//
// Prints the CLI help message with all available commands
func (c *CLI) help() {
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
func (c *CLI) loop() error {

	// print current logged in user before prompt
	if c.user != nil {
		fmt.Printf("%s ", c.user.email)
	}
	fmt.Print("-> ")

	text, err := c.reader.ReadString('\n')
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
		c.help()
	case "quit":
		c.running = false
	case "get":
		if len(words) != 2 {
			c.errorLog.Println("Please provide a stock symbol to get")
			break
		}

		var result string
		result, err = c.stocksClient.GetQuote(words[1])
		if err == nil {
			fmt.Println(result)
		}
	case "create":
		err = c.createUser()
	case "login":
		err = c.login()
	case "user":
		err = c.getUser()
	case "logout":
		c.logout()
	default:
		c.errorLog.Println("Invalid command, see reference for command")
	}

	return err
}

// Gets information of the current logged in user, or prints an error if no
// user is currently logged in
func (c *CLI) getUser() error {
	if c.user == nil {
		c.errorLog.Printf("Error: must be logged in to get user info")
		return nil
	}

	user, err := c.stocksClient.GetUser(c.user.email, c.user.accessToken)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.Unauthenticated {
			c.errorLog.Print(err)
			return nil
		}
		return err
	}

	fmt.Printf("Username: %s, Email: %s, First Name: %s, Last Name: %s",
		user.Username, user.Email, user.FirstName, user.LastName)
	return nil
}

// Logs out the current user or prints an error message if no user is logged in
func (c *CLI) logout() {
	if c.user != nil {
		c.user = nil
	} else {
		c.errorLog.Printf("Error: attempted to log out when no user is logged in")
	}
}

// Logs into a user account by prompting for email and password and making
// request to the stocksClient
func (c *CLI) login() error {
	if c.user != nil {
		c.errorLog.Print("Error: already logged in, to log in again please log out first")
		return nil
	}

	fmt.Printf("Enter email: ")
	text, err := c.reader.ReadString('\n')
	if err != nil {
		return nil
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		c.errorLog.Print("no email entered")
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
	token, err := c.stocksClient.LoginUser(email, string(bytePassword))
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.Unauthenticated {
			c.errorLog.Print("Invalid credentials")
			return nil
		}
		return err
	}

	c.user = &authenticatedUser{email: email, accessToken: token}
	fmt.Println()
	return nil
}

// createUser gets the appropriate fields from reader to create a user on
// the stocks client
func (c *CLI) createUser() error {
	fmt.Printf("Enter the following fields separated by spaces: firstname, lastname, username, email, password:\n-> ")
	text, err := c.reader.ReadString('\n')
	if err != nil {
		return err
	}

	words := strings.Fields(text)
	if len(words) != 5 {
		c.errorLog.Printf("Error: incorrect number of fields for create user")
		return nil
	}

	id, err := c.stocksClient.CreateUser(words[0], words[1], words[2], words[3], words[4])

	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			c.errorLog.Print("user already exists")
			return nil
		} else {
			return err
		}
	}

	fmt.Printf("user created with id: %d\n", id)
	return nil
}
