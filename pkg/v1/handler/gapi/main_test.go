package gapi

import (
	"log"
	"os"
	"testing"

	"github.com/colin-mcl/stocks/internal/db"
	"github.com/colin-mcl/stocks/internal/token"
	"github.com/colin-mcl/stocks/pkg/v1/repo"
	"github.com/colin-mcl/stocks/pkg/v1/usecase"
	"github.com/colin-mcl/stocks/util"
)

// test server uses real DB connection
// TODO: mock this? might not be worth the trouble
var testServer *Server

// Runs all *_test.go files found in the package
func TestMain(m *testing.M) {
	db, err := db.NewDBConn()
	if err != nil {
		panic(err)
	}

	maker, err := token.NewPasetoMaker(util.RandomString(32))
	if err != nil {
		panic(err)
	}

	testServer = &Server{
		api_key:    os.Getenv("STOCKS_API_KEY"),
		infoLog:    log.New(os.Stdout, "INFO ", log.Ldate),
		errorLog:   log.New(os.Stderr, "ERROR ", log.Ldate),
		uc:         usecase.NewUC(repo.NewRepo(db)),
		tokenMaker: maker,
	}

	os.Exit(m.Run())
}
