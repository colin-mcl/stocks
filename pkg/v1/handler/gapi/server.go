package gapi

// grpc Server implementation

import (
	"database/sql"
	"log"
	"os"

	"github.com/colin-mcl/stocks/internal/token"
	"github.com/colin-mcl/stocks/pb"
	"github.com/colin-mcl/stocks/pkg/v1/repo"
	"github.com/colin-mcl/stocks/pkg/v1/usecase"
)

// Server serves gRPC requests for our stocks service
type Server struct {
	// Allows forwards compatibility as the unimplemented server can accept all gRPC requests before they are implemented
	pb.UnimplementedStocksServer

	// UseCase allows for CRUD operations on users and positions
	uc usecase.UseCaseInterface

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
	infoLog *log.Logger,
	maker token.Maker) *Server {

	server := &Server{
		uc:         usecase.NewUC(repo.NewRepo(db)),
		tokenMaker: maker,
		api_key:    os.Getenv("STOCKS_API_KEY"),
		errorLog:   errorLog,
		infoLog:    infoLog,
	}

	return server
}
