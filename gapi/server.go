package gapi

// grpc Server implementation

import (
	"database/sql"
	"log"
	"os"

	"github.com/colin-mcl/stocks/pb"
)

// Server serves gRPC requests for our stocks service
type Server struct {
	// Allows forwards compatibility as the unimplemented server can accept all gRPC requests before they are implemented
	pb.UnimplementedStocksServer

	// stores a pool of db connections for operations on our stocks database
	db *sql.DB

	// api key required for the Yahoo finance api
	// set the value with STOCKS_API_KEY env variable
	api_key string

	// Custom logger object to print error log messages
	errorLog *log.Logger

	// Custom logger object to print info level log messages
	infoLog *log.Logger
}

func NewServer(db *sql.DB, errorLog *log.Logger, infoLog *log.Logger) (*Server, error) {
	server := &Server{}
	// server.api_key = "4XKTWpU6YY2Y3N6zGKdip6iICRouIJmM83ePOUWD"
	server.db = db
	server.api_key = os.Getenv("STOCKS_API_KEY")
	server.errorLog = errorLog
	server.infoLog = infoLog

	return server, nil
}
