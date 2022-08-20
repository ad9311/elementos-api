package db

import (
	"context"
	"net/http"
	"time"
)

// GetUser queries a given user and returns it.
func (d *Database) GetUser(r *http.Request) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user := User{}
	var storedPassword string
	query := "select * from users where username = $1"
	row := d.Conn.QueryRowContext(ctx, query, r.FormValue("username"))
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&storedPassword,
		&user.DefaultUser,
		&user.LastLogin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}
