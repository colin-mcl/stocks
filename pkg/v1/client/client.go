package client

import (
	"context"
	"strings"
	"time"

	"github.com/colin-mcl/stocks/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

// Structure:

// GetQuote: requires NO authorization, users do not need to be logged in
// LoginUser: user passes their email and password and is returned an access
//			  token
//
//

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
func (stocksClient *StocksClient) GetQuote(symbol string) (string, error) {
	// Create RPC request
	req := &pb.GetQuoteRequest{Symbol: strings.ToUpper(symbol)}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := stocksClient.service.GetQuote(ctx, req)
	if err != nil {
		return "", err
	}

	return protojson.Format(resp), nil
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

// LoginUser calls the LoginUser service on the stocks service with the
// email and password and returns the access token if the login is suceessful
func (stocksClient *StocksClient) LoginUser(email string, password string) (string, error) {
	req := &pb.LoginUserRequest{Email: email, Password: password}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := stocksClient.service.LoginUser(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.GetAccessToken(), nil
}

func (stocksClient *StocksClient) GetUser(email, accessToken string) (*pb.User, error) {
	req := &pb.GetUserRequest{Email: email}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = metadata.AppendToOutgoingContext(ctx, "authentication", accessToken)

	resp, err := stocksClient.service.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.GetUser(), nil
}
