package gapi_test

import (
	"context"
	"testing"

	"github.com/colin-mcl/stocks/pb"
	"github.com/colin-mcl/stocks/util"
	"github.com/stretchr/testify/assert"
)

func TestGetQuote(t *testing.T) {
	type cases struct {
		name  string
		input *pb.GetQuoteRequest
	}

	for _, scenario := range []cases{
		{
			name:  "bad symbol short",
			input: &pb.GetQuoteRequest{Symbol: "1"},
		},
		{
			name:  "bad symbol long random",
			input: &pb.GetQuoteRequest{Symbol: util.RandomString(255)},
		},
		{
			name:  "empty symbol",
			input: &pb.GetQuoteRequest{Symbol: ""},
		},
		{
			name:  "tsla1",
			input: &pb.GetQuoteRequest{Symbol: "tsla1"},
		}} {
		t.Run(scenario.name, func(t *testing.T) {
			resp, err := testServer.GetQuote(context.Background(), scenario.input)
			assert.Error(t, err)
			assert.Nil(t, resp)
			assert.ErrorContains(t, err, "bad symbol")
		})
	}
	for _, scenario := range []cases{
		{
			name:  "tsla",
			input: &pb.GetQuoteRequest{Symbol: "TSLA"},
		},
		{
			name:  "aapl",
			input: &pb.GetQuoteRequest{Symbol: "AAPL"},
		},
		{
			name:  "fxaix",
			input: &pb.GetQuoteRequest{Symbol: "FXAIX"},
		}} {
		t.Run(scenario.name, func(t *testing.T) {
			resp, err := testServer.GetQuote(context.Background(), scenario.input)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, scenario.input.GetSymbol(), resp.Quote.GetSymbol())
			assert.NotZero(t, resp.Quote.GetRegularMarketPrice())
			assert.Equal(t, "USD", resp.Quote.GetCurrency())
		})
	}
}
