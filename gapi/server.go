package gapi

// grpc Server implementation

import (
	"database/sql"
	"log"
	"os"

	"github.com/colin-mcl/stocks/internal/models"
	"github.com/colin-mcl/stocks/pb"
	"github.com/colin-mcl/stocks/token"
	"github.com/colin-mcl/stocks/util"
)

// Server serves gRPC requests for our stocks service
type Server struct {
	// Allows forwards compatibility as the unimplemented server can accept all gRPC requests before they are implemented
	pb.UnimplementedStocksServer

	// Allows the user  objects to be available on the GRPC server
	users *models.UserModel

	// TokenMaker for making and verifying user authorization tokens
	tokenMaker token.Maker

	// api key required for the Yahoo finance api
	// set the value with STOCKS_API_KEY env variable
	api_key string

	// Custom logger object to print error log messages
	errorLog *log.Logger

	// Custom logger object to print info level log messages
	infoLog *log.Logger
}

func NewServer(
	db *sql.DB,
	errorLog *log.Logger,
	infoLog *log.Logger) (*Server, error) {

	// Create a PasetoMaker with symmetric key
	// TODO: make key an env variable
	maker, err := token.NewPasetoMaker(util.RandomString(32))
	if err != nil {
		return nil, err
	}

	server := &Server{
		users:      &models.UserModel{DB: db},
		tokenMaker: maker,
		api_key:    os.Getenv("STOCKS_API_KEY"),
		errorLog:   errorLog,
		infoLog:    infoLog,
	}
	// server.api_key = "4XKTWpU6YY2Y3N6zGKdip6iICRouIJmM83ePOUWD"

	return server, nil
}
