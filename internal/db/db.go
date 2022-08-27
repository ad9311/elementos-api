package db

import (
	"database/sql"
	"strings"
	"time"

	// Drivers for PostgreSQL
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Database ...
type Database struct {
	Conn *sql.DB
}

const (
	maxOpenConn = 10
	maxIdleConn = 5
	maxLifeTime = 5 * time.Minute
)

var database Database

// New ...
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

func pgArrayToSlice(pgArr string) []string {
	chars := []string{"{", "}"}
	for _, v := range chars {
		pgArr = strings.ReplaceAll(pgArr, v, "")
	}
	slice := strings.Split(pgArr, ",")

	return slice
}
