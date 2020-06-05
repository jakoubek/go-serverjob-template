package main

import (
	"database/sql"

	// ODBC-Zugriff auf die ZSP
	_ "github.com/alexbrainman/odbc"
)

// DB stands for the database connection
type DB struct {
	Db *sql.DB
}

// Open opens a database connection and returns it
func NewDbConnection(dsn string) (*DB, error) {
	Db, err := sql.Open("odbc", dsn)
	if err != nil {
		return nil, err
	}
	return &DB{Db}, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.Db.Close()
}
