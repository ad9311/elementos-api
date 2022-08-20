package db

import (
	"database/sql"
	"time"

	// Drivers for PostgreSQL
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Database is used to contain a database connection.
type Database struct {
	Conn *sql.DB
}

const (
	maxOpenConn = 10
	maxIdleConn = 5
	maxLifeTime = 5 * time.Minute
)

var database Database

// New connects to a database and resturns it as a connection.
func New(dsn string) (*Database, error) {
	db, err := sql.Open("pgx", dsn)
	database.Conn = db
	if err != nil {
		return &database, err
	}

	if err = db.Ping(); err != nil {
		return &database, err
	}

	database.Conn.SetMaxOpenConns(maxOpenConn)
	database.Conn.SetMaxIdleConns(maxIdleConn)
	database.Conn.SetConnMaxLifetime(maxLifeTime)

	return &database, nil
}
