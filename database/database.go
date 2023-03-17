package database

import (
	"database/sql"
)

type Database struct {
	conn *sql.DB
}

func (s Database) Close() error {
	return s.conn.Close()
}
