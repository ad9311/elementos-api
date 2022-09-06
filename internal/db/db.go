package db

import (
	"database/sql"
	"fmt"
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
	chars := []string{"{", "}", `"`}
	for _, v := range chars {
		pgArr = strings.ReplaceAll(pgArr, v, "")
	}
	slice := strings.Split(pgArr, ",")

	return slice
}

func parseLandmarkQueries(baseQuery string, urlQueries map[string]string) (string, error) {
	first := true
	for k, v := range urlQueries {
		if strings.Contains(k, "sel_") {
			if first {
				baseQuery += " WHERE "
				first = false
			} else {
				baseQuery += " AND "
			}

			if strings.Contains(k, "sel_arr_") {
				baseQuery += fmt.Sprintf("'%s'=ANY(landmarks.%s)", v, strings.Split(k, "sel_arr_")[1])
			} else {
				baseQuery += fmt.Sprintf("landmarks.%s='%s'", strings.Split(k, "sel_")[1], v)
			}
		}
	}

	if v, ok := urlQueries["ord_order_by"]; ok {
		baseQuery += fmt.Sprintf(" ORDER BY landmarks.%s", v)
		if _, ok := urlQueries["ord_desc"]; ok {
			baseQuery += " DESC"
		}
	} else {
		if _, ok := urlQueries["ord_desc"]; ok {
			return baseQuery, fmt.Errorf("order_by baseQuery is missing")
		}
	}

	return baseQuery + ";", nil
}
