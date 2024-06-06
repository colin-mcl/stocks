package models

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testModels *UserModel

// Runs all *_test.go files found in the package
func TestMain(m *testing.M) {
	testDB, err := sql.Open("mysql", "web:Amsterdam22!@/stocks?parseTime=true")

	testModels = &UserModel{DB: testDB}
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
