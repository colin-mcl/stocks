package gapi

import (
	"context"
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

func TestGetQuoteBadKey(t *testing.T) {
	server := &Server{api_key: "bad"}
	resp, err := server.GetQuote(context.Background(), &pb.GetQuoteRequest{Symbol: "TSLA"})

	require.Nil(t, resp)
	require.Error(t, err)
	require.Equal(t, err, status.Errorf(codes.Internal, "bad API key: bad"))
}

func TestGetQuote(t *testing.T) {
	server := &Server{api_key: good_key}

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
