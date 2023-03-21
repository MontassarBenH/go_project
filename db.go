package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "tasks.db")
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}
	return db, nil
}
