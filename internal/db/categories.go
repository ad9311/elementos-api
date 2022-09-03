package db

import (
	"context"
	"time"
)

// SelectCategories ...
func (d *Database) SelectCategories() ([]Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	categories := []Category{}
	category := Category{}
	query := `SELECT * FROM categories ORDER BY name;`

	rows, err := d.Conn.QueryContext(ctx, query)
	if err != nil {
		rows.Close()
		return categories, err
	}

	for rows.Next() {
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			rows.Close()
			return categories, err
		}
		categories = append(categories, category)
	}
	rows.Close()

	return categories, nil
}

// SelectCategory ...
func (d *Database) SelectCategory(id int64) (Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	category := Category{}
	query := `SELECT * FROM categories WHERE id=$1`

	row := d.Conn.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		return category, err
	}

	return category, nil
}

// InsertCategory ...
func (d *Database) InsertCategory(formMap map[string]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO categories
	("name",created_at,updated_at)
	values ($1,$2,$3);
	`
	_, err := d.Conn.ExecContext(
		ctx,
		query,
		formMap["name"],
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}

// UpdateCategory ...
func (d *Database) UpdateCategory(formMap map[string]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE categories SET name=$1,updated_at=$2 WHERE id=$3;`

	_, err := d.Conn.ExecContext(
		ctx,
		query,
		formMap["name"],
		time.Now(),
		formMap["category_id"],
	)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCategory ...
func (d *Database) DeleteCategory(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM categories WHERE id=$1;"

	_, err := d.Conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
