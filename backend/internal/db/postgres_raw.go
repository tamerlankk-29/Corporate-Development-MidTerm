package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(time.Hour)
	return db, nil
}
