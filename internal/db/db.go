package db

import (
	"database/sql"
	"time"

	// Drivers for PostgreSQL
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	maxOpenConn = 10
	maxIdleConn = 5
	maxLifeTime = 5 * time.Minute
)

// New connects to a database and resturns it as a connection.
func New(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return db, err
	}

	if err = db.Ping(); err != nil {
		return db, err
	}

	db.SetMaxOpenConns(maxOpenConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(maxLifeTime)

	return db, nil
}
