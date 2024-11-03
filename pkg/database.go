package pkg

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDbConn() (*sql.DB, error) {
	filePath := "./db.sqlite"
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return nil, err
	}

	return db, nil
}
