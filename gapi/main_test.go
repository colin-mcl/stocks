package gapi

import (
	"log"
	"os"
	"testing"

	"github.com/colin-mcl/stocks/internal/models"
	"github.com/colin-mcl/stocks/util"
)

// Runs all *_test.go files found in the package
func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

func makeDefaultServer() *Server {
	db, err := util.OpenDB("web:Amsterdam22!@/stocks?parseTime=true")
	if err != nil {
		panic(err)
	}

	server := &Server{
		api_key:  "4XKTWpU6YY2Y3N6zGKdip6iICRouIJmM83ePOUWD",
		infoLog:  log.New(os.Stdout, "INFO ", log.Ldate),
		errorLog: log.New(os.Stderr, "ERROR ", log.Ldate),
		users:    &models.UserModel{DB: db},
	}

	return server
}
