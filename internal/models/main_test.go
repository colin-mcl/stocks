package models

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var users *UserModel
var positions *PositionModel

// Runs all *_test.go files found in the package
func TestMain(m *testing.M) {
	testDB, err := sql.Open("mysql", "web:Amsterdam22!@/stocks?parseTime=true&loc=America%2FNew_York")

	users = &UserModel{DB: testDB}
	positions = &PositionModel{DB: testDB}
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
