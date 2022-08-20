package db

import (
	"context"
	"net/http"
	"time"
)

// SelectUser queries a given user and returns it.
func (d *Database) SelectUser(r *http.Request) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user := User{}
	query := "SELECT * FROM users WHERE username = $1;"
	row := d.Conn.QueryRowContext(ctx, query, r.PostFormValue("username"))
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.EncryptedPassword,
		&user.DefaultUser,
		&user.LastLogin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

// UpdateLastLogin updates the last_login column in the users table
// each time a userssigns in.
func (d *Database) UpdateLastLogin(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "UPDATE users SET last_login=$1 WHERE id=$2;"
	date := time.Now()
	_, err := d.Conn.ExecContext(ctx, query, date, user.ID)
	if err != nil {
		return err
	}
	user.LastLogin = date

	return nil
}

// InsertUser inserts a new user in the database
func (d *Database) InsertUser(r *http.Request, encryptedPassword string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO users
	(first_name,last_name,username,email,password,default_user,
	last_login,created_at,updated_at)
	values ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`
	_, err := d.Conn.ExecContext(
		ctx,
		query,
		r.PostFormValue("first_name"),
		r.PostFormValue("last_name"),
		r.PostFormValue("username"),
		r.PostFormValue("email"),
		encryptedPassword,
		false,
		time.Now(),
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
