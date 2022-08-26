package db

import (
	"context"
	"time"
)

// SelectInvitation ...
func (d *Database) SelectInvitation(code string) (Inviation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	inviation := Inviation{}
	query := "SELECT * FROM invitations WHERE code=$1;"
	row := d.Conn.QueryRowContext(ctx, query, code)
	err := row.Scan(
		&inviation.ID,
		&inviation.Code,
		&inviation.ExpiresAt,
		&inviation.CreatedAt,
		&inviation.UpdatedAt,
	)
	if err != nil {
		return inviation, err
	}

	return inviation, nil
}
