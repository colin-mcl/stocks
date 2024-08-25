package client

import (
	"context"
	"strings"
	"time"

	"github.com/colin-mcl/stocks/internal/models"
	"github.com/colin-mcl/stocks/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

// Client
//
// The client package wraps the GRPC API used to communicate with the stocks
// server, providing a simple interface that accepts golang standard types
// and returns the expected results

type StocksClient struct {
	// Internal StocksClient service
	service pb.StocksClient
}

// NewStocksClient
//
// Returns a new StocksClient instance
// A grpc client connection must be provided to communicate with the server
func NewStocksClient(conn *grpc.ClientConn) *StocksClient {

	return &StocksClient{
		service: pb.NewStocksClient(conn)}
}

// GetQuote
//
// Given a stock symbol makes the GRPC GetQuote to the server and returns the
// results in a formatted string
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

// CreateUser
//
// Given the fields of a user to create makes the GRPC CreateUser to the server
// and returns the id of the new user if successful
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

// LoginUser
//
// calls the LoginUser service on the stocks service with the
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

// GetUser
//
// Makes the GRPC GetUser to the server given the provided email and access token
// returns the result as a models.User pointer if successful
func (stocksClient *StocksClient) GetUser(email, accessToken string) (*models.User, error) {
	req := &pb.GetUserRequest{Email: email}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = metadata.AppendToOutgoingContext(ctx, "authentication", accessToken)

	resp, err := stocksClient.service.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return &models.User{
		Username:  resp.GetUser().GetUsername(),
		Email:     resp.GetUser().GetEmail(),
		FirstName: resp.GetUser().GetFirstName(),
		LastName:  resp.GetUser().GetLastName(),
	}, nil
}
