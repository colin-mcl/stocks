package gapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/colin-mcl/stocks/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// TODO: make region and language options
const (
	yahooURL = "https://yfapi.net/v6/finance/quote?symbols=%s&region=US&lang=en"
)

// badResponse indicates that the api key provided (or lack thereof) did not
// successfully work with the yahoo finance api and should be replaced
type badResponse struct {
	Message string `json:"message"`
	Hint    string `json:"hint"`
}

// Example request:
// https://yfapi.net/v6/finance/quote?region=US&lang=en&symbols=AAPL
func (server *Server) GetTicker(ctx context.Context, r *pb.GetTickerRequest) (*pb.GetTickerResponse, error) {
	if api_key == "" {
		initKey()
	}

	// Create new HTTP request and add API key to the header
	req, err := initRequest("GET", fmt.Sprintf(yahooURL, r.GetSymbol()), nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create GET request")
	}

	// Make HTTP request with the default client
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to execute get request, %s", err)

	}
	defer response.Body.Close()

	// Read the body of the response as a slice of bytes and reset the io Reader
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read body of request")
	}

	response.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var badKey badResponse

	// Check if we got a bad response bc of no API key set
	d := json.NewDecoder(response.Body)
	err = d.Decode(&badKey)
	if err != nil || badKey.Message != "" {
		return nil, status.Errorf(codes.Internal, "bad API key")
	}

	y := &pb.YhResponse{}
	err = protojson.Unmarshal(bodyBytes, y)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmarshal to yahoo response: %s", err)
	}

	// q := &pb.Quote{}
	// err = protojson.Unmarshal(y.GetQ(), q)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "failed to unmarshal to quote : %s", err)
	// }

	// t := &pb.Ticker{}
	// err = protojson.Unmarshal(q.GetResult(), t)

	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "failed to unmarshal to ticker: %s", err)
	// }

	return &pb.GetTickerResponse{Ticker: y.GetQuoteResponse().GetResult()[0]}, nil
}
