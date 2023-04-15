package database

import "database/sql"

type ClientDB struct {
	DB *sql.DB
}

func NewClient(db *sql.DB) *ClientDB {
	return &ClientDB{
		DB: db,
	}
}
