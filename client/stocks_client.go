package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/colin-mcl/stocks/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

// Desired structure:

// GetQuote: requires NO authorization, users do not need to be logged in
// LoginUser: user passes their email and password and is returned an access
//			  token which is stored WHERE???
//
// AuthorizedRPC: either stored access token is added to the context or
//				   LoginUser is called to get the access token before the
//				   authorizedRPC is called

// Idea:
// map [string] string authorizedUsers: email -> accessToken
// if the user is not found in the authorizedUsers, tell them to call login
// and try again
//
// on Login: add user email and access token to the map, display users
// 	   		 email on the terminal as visual display

// StocksClient provides a simple wrapper for the stocks service rpcs
type StocksClient struct {
	// Internal StocksClient service
	service pb.StocksClient
}

// Creates a Stocks client with the provided grpc client connection
func NewStocksClient(conn *grpc.ClientConn) *StocksClient {

	return &StocksClient{
		service: pb.NewStocksClient(conn)}
}

// Gets and prints a simple stock quote
func (stocksClient *StocksClient) GetQuote(symbol string) error {
	// Create RPC request
	req := &pb.GetQuoteRequest{Symbol: strings.ToUpper(symbol)}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := stocksClient.service.GetQuote(ctx, req)
	if err != nil {
		return err
	}

	fmt.Println(protojson.Format(resp))
	return nil
}

func (stocksClient *StocksClient) CreateUser(
	firstName string,
	lastName string,
	username string,
	email string,
	password string) (int, error) {

	req := &pb.CreateUserRequest{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := stocksClient.service.CreateUser(ctx, req)
	if err != nil {
		return -1, err
	}

	return int(resp.GetId()), nil
}
