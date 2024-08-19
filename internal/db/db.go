package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDBConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "web:Amsterdam22!@/stocks?parseTime=true")

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
