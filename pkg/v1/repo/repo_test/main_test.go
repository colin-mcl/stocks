package repo_test

import (
	"log"
	"os"
	"testing"

	"github.com/colin-mcl/stocks/internal/db"
	"github.com/colin-mcl/stocks/pkg/v1/repo"
)

var testRepo repo.RepoInterface

// test all repository interface functionality
// currently uses real DB for testing
// change to mock???
func TestMain(m *testing.M) {
	db, err := db.NewDBConn()

	if err != nil {
		log.Fatal(err)
	}
	testRepo = repo.NewRepo(db)

	os.Exit(m.Run())
}
