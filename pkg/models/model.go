package models

import "database/sql"

type Model struct {
	Conn *sql.DB
}
