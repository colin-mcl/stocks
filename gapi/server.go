package gapi

import (
	"fmt"
	"os"

	"github.com/colin-mcl/stocks/pb"
)

// Server serves gRPC requests for our stocks service
type Server struct {
	// Allows forwards compatibility as the unimplemented server can accept all gRPC requests before they are implemented
	pb.UnimplementedStocksServer
	api_key string
}

func NewServer() (*Server, error) {
	server := &Server{}
	server.api_key = os.Getenv("STOCKS_API_KEY")
	fmt.Printf("api_key = %s\n", server.api_key)

	return server, nil
}
