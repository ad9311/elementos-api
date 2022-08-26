package db

import (
	"context"
	"time"
)

// SelectUserByUsername ...
func (d *Database) SelectUserByUsername(username string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user := User{}
	query := `SELECT id,first_name,last_name,username,password
	FROM users WHERE username=$1;`

	row := d.Conn.QueryRowContext(ctx, query, username)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.HashedPassword,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}

// UpdateUserLastLogin ...
func (d *Database) UpdateUserLastLogin(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "UPDATE users SET last_login=$1 WHERE id=$2;"
	date := time.Now()

	_, err := d.Conn.ExecContext(ctx, query, date, id)
	if err != nil {
		return err
	}

	return nil
}

// InsertUser ...
func (d *Database) InsertUser(formMap map[string]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO users
	(first_name,last_name,username,password,
	last_login,created_at,updated_at)
	values ($1,$2,$3,$4,$5,$6,$7,$8,$9);
	`
	_, err := d.Conn.ExecContext(
		ctx,
		query,
		formMap["first_name"],
		formMap["last_name"],
		formMap["username"],
		formMap["password"],
		time.Now(),
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}
