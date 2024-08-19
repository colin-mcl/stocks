package repo

import (
	"log"
	"os"
	"testing"

	db "github.com/colin-mcl/stocks/internal/db"
)

var testRepo RepoInterface

// test all repository interface functionality
// currently uses real DB for testing
// change to mock???
func TestMain(m *testing.M) {
	db, err := db.NewDBConn()

	if err != nil {
		log.Fatal(err)
	}
	testRepo = NewRepo(db)

	os.Exit(m.Run())
}
