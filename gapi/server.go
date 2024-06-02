package gapi

// grpc Server implementation

import (
	"github.com/colin-mcl/stocks/pb"
)

// Server serves gRPC requests for our stocks service
type Server struct {
	// Allows forwards compatibility as the unimplemented server can accept all gRPC requests before they are implemented
	pb.UnimplementedStocksServer

	// api key required for the Yahoo finance api
	// set the value with STOCKS_API_KEY env variable
	api_key string
}

func NewServer() (*Server, error) {
	server := &Server{}
	server.api_key = "4XKTWpU6YY2Y3N6zGKdip6iICRouIJmM83ePOUWD"
	// server.api_key = os.Getenv("STOCKS_API_KEY")

	return server, nil
}
