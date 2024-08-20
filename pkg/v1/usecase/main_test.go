package usecase

import (
	"log"
	"os"
	"testing"

	db "github.com/colin-mcl/stocks/internal/db"
	"github.com/colin-mcl/stocks/pkg/v1/repo"
)

var testUC UseCaseInterface

func TestMain(m *testing.M) {
	db, err := db.NewDBConn()

	if err != nil {
		log.Fatal(err)
	}

	testUC = NewUC(repo.NewRepo(db))
	os.Exit(m.Run())
}
