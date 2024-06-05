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
// 4XKTWpU6YY2Y3N6zGKdip6iICRouIJmM83ePOUWD - API key
const (
	yahooURL = "https://yfapi.net/v6/finance/quote?symbols=%s&region=US&lang=en"
)

// badResponse indicates that the api key provided (or lack thereof) did not
// successfully work with the yahoo finance api and should be replaced
type badResponse struct {
	Message string `json:"message"`
	Hint    string `json:"hint"`
}

// Handles the GetQuote grpc call by making a get request to the Yahoo
// finance API and returning the result as a quote object
// Example request:
// https://yfapi.net/v6/finance/quote?region=US&lang=en&symbols=AAPL
func (server *Server) GetQuote(ctx context.Context, r *pb.GetQuoteRequest) (*pb.GetQuoteResponse, error) {
	server.infoLog.Printf("get quote request recieved: %s\n", r.GetSymbol())

	// Create new HTTP request and add API key to the header
	req, err := http.NewRequest("GET", fmt.Sprintf(yahooURL, r.GetSymbol()), nil)
	req.Header.Set("x-api-key", server.api_key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create request %s", err)
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
		return nil, status.Errorf(codes.Internal, "bad API key: %s", server.api_key)
	}

	y := &pb.YhResponse{}
	err = protojson.Unmarshal(bodyBytes, y)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmarshal to yahoo response: %s", err)
	}

	// Check if we got 0 quotes in the result, implying an invalid symbol
	quotes := y.GetQuoteResponse().GetResult()
	if len(quotes) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid quote symbol: %s", r.GetSymbol())
	}

	return &pb.GetQuoteResponse{Quote: quotes[0]}, nil
}
