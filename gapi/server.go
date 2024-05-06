package gapi

import "github.com/colin-mcl/stocks/pb"

// Server serves gRPC requests for our stocks service
type Server struct {
	// Allows forwards compatibility as the unimplemented server can accept all gRPC requests before they are implemented
	pb.UnimplementedStocksServer
}

func NewServer() (*Server, error) {
	server := &Server{}

	return server, nil
}
