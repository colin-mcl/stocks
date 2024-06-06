package gapi

import (
	"bufio"
	"context"
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/colin-mcl/stocks/pb"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/* This file tests the functionality of the GetQuote rpc function wihtout
 * creating a connection.
 */

// Working API key
const good_key = "4XKTWpU6YY2Y3N6zGKdip6iICRouIJmM83ePOUWD"

// Tests the response of GetQuote if the server has a bad api key
func TestGetQuoteBadKey(t *testing.T) {
	server := makeDefaultServer()
	server.api_key = "bad"
	resp, err := server.GetQuote(context.Background(), &pb.GetQuoteRequest{Symbol: "TSLA"})

	require.Nil(t, resp)
	require.Error(t, err)
	require.Equal(t, err, status.Errorf(codes.Internal, "bad API key: bad"))
}

func TestGetQuote(t *testing.T) {
	server := makeDefaultServer()

	tests := []struct {
		name string
		req  *pb.GetQuoteRequest
		err  error
	}{
		{
			name: "TSLA",
			req:  &pb.GetQuoteRequest{Symbol: "TSLA"},
			err:  nil,
		},
		{
			name: "BADSYMBOL",
			req:  &pb.GetQuoteRequest{Symbol: "BADSYMBOL"},
			err:  status.Errorf(codes.InvalidArgument, "invalid quote symbol: BADSYMBOL"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := server.GetQuote(context.Background(), tt.req)

			if resp != nil {
				require.Nil(t, err)
				require.Equal(t, tt.name, resp.GetQuote().Symbol)
			} else {
				require.Nil(t, resp)
				require.NotNil(t, err)
				require.Equal(t, err, tt.err)
			}
		})
	}
}

// Tests the GetQuote RPC using a random selection of
// valid NYSE symbols, expected to be stored in the file symbols.txt one
// level up froim the gapi directory.
func TestGetQuoteAll(t *testing.T) {
	file, err := os.Open("../symbols.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	server := makeDefaultServer()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		symbol := scanner.Text()

		// only test this symbol if rand() generates 999a
		i := rand.Intn(1000)
		if i != 999 {
			continue
		}
		resp, err := server.GetQuote(context.Background(),
			&pb.GetQuoteRequest{Symbol: symbol})

		if err != nil {
			log.Fatal(err)
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, symbol, resp.GetQuote().Symbol)
	}
}
