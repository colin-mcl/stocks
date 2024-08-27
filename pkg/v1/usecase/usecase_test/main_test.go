package usecase_test

import (
	"log"
	"os"
	"testing"

	db "github.com/colin-mcl/stocks/internal/db"
	"github.com/colin-mcl/stocks/pkg/v1/repo"
	"github.com/colin-mcl/stocks/pkg/v1/usecase"
)

var testUC usecase.UseCaseInterface

func TestMain(m *testing.M) {
	db, err := db.NewDBConn()

	if err != nil {
		log.Fatal(err)
	}

	testUC = usecase.NewUC(repo.NewRepo(db))
	os.Exit(m.Run())
}
