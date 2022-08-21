package db

import (
	"context"
	"time"
)

// SelectInvitationCode queries a given invitation code
// and returns how many times it has been used.
func (d *Database) SelectInvitationCode(code string) (InviationCodes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ic := InviationCodes{}
	query := "SELECT * FROM invitation_codes WHERE code=$1;"
	row := d.Conn.QueryRowContext(ctx, query, code)
	err := row.Scan(
		&ic.ID,
		&ic.Code,
		&ic.Validity,
		&ic.CreatedAt,
		&ic.UpdatedAt,
	)
	if err != nil {
		return ic, err
	}

	return ic, nil
}
