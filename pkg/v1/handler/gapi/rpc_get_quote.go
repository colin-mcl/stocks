package gapi

import (
	"context"

	"github.com/colin-mcl/stocks/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handles the GetQuote grpc call by calling usecase GetRequest
func (server *Server) GetQuote(ctx context.Context, r *pb.GetQuoteRequest) (*pb.GetQuoteResponse, error) {
	server.infoLog.Printf("get quote request recieved: %s\n", r.GetSymbol())

	if r.GetSymbol() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "bad symbol")
	}

	quote, err := server.uc.GetQuote(r.GetSymbol())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error: failed to get quote %s", err)
	}

	return &pb.GetQuoteResponse{Quote: &pb.Quote{
		Symbol:             quote.Symbol,
		Region:             quote.Region,
		ShortName:          quote.ShortName,
		TimezoneShort:      quote.TimezoneShort,
		Market:             quote.Market,
		Currency:           quote.Currency,
		FiftyTwoWeekLow:    quote.FiftyTwoWeekLow,
		FiftyTwoWeekHigh:   quote.FiftyTwoWeekHigh,
		RegularMarketPrice: quote.RegularMarketPrice,
	}}, nil
}
