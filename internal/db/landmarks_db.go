package db

import (
	"context"
	"time"
)

// InsertLandmark ...
func (d *Database) InsertLandmark(formMap map[string]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO landmarks
	(name,native_name,description,wiki_url,
	location,img_urls,user_id,category_id,created_at,updated_at)
	values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);
	`
	_, err := d.Conn.ExecContext(
		ctx,
		query,
		formMap["name"],
		formMap["native_name"],
		formMap["description"],
		formMap["wiki_url"],
		formMap["location"],
		formMap["img_urls"],
		formMap["user_id"],
		formMap["category_id"],
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}

// SelectLandmarks ...
func (d *Database) SelectLandmarks() ([]Landmark, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	landmarks := []Landmark{}
	landmark := Landmark{}
	location := ""
	imgURLs := ""
	query := `SELECT landmarks.*,users.username,categories.name FROM landmarks
	LEFT JOIN users ON users.id=landmarks.user_id
	LEFT JOIN categories ON categories.id=landmarks.category_id ORDER BY landmarks.name;
	`

	rows, err := d.Conn.QueryContext(ctx, query)
	if err != nil {
		rows.Close()
		return landmarks, err
	}

	for rows.Next() {
		err := rows.Scan(
			&landmark.ID,
			&landmark.Name,
			&landmark.NativeName,
			&landmark.Description,
			&landmark.WikiURL,
			&location,
			&imgURLs,
			&landmark.Default,
			&landmark.UserID,
			&landmark.CreatedAt,
			&landmark.UpdatedAt,
			&landmark.CategoryID,
			&landmark.CreatedBy,
			&landmark.Category,
		)
		if err != nil {
			rows.Close()
			return landmarks, err
		}

		landmark.Location = pgArrayToSlice(location)
		landmark.ImgURLs = pgArrayToSlice(imgURLs)
		landmarks = append(landmarks, landmark)
	}
	rows.Close()

	return landmarks, nil
}

// SelectLandmarkByID ...
func (d *Database) SelectLandmarkByID(id int64) (Landmark, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	landmark := Landmark{}
	location := ""
	imgURLs := ""
	query := `SELECT landmarks.*,users.username,categories.name FROM landmarks
	LEFT JOIN users ON users.id=landmarks.user_id
	LEFT JOIN categories ON categories.id=landmarks.category_id
	WHERE landmarks.id=$1;
	`

	row := d.Conn.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&landmark.ID,
		&landmark.Name,
		&landmark.NativeName,
		&landmark.Description,
		&landmark.WikiURL,
		&location,
		&imgURLs,
		&landmark.Default,
		&landmark.UserID,
		&landmark.CreatedAt,
		&landmark.UpdatedAt,
		&landmark.CategoryID,
		&landmark.CreatedBy,
		&landmark.Category,
	)
	if err != nil {
		return landmark, err
	}

	landmark.Location = pgArrayToSlice(location)
	landmark.ImgURLs = pgArrayToSlice(imgURLs)

	return landmark, nil
}

// SelectLandmarkByName ...
func (d *Database) SelectLandmarkByName(name string) (Landmark, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	landmark := Landmark{}
	location := ""
	imgURLs := ""
	query := `SELECT landmarks.*,users.username,categories.name FROM landmarks
	LEFT JOIN users ON users.id=landmarks.user_id
	LEFT JOIN categories ON categories.id=landmarks.category_id
	WHERE landmarks.name=$1;
	`

	row := d.Conn.QueryRowContext(ctx, query, name)
	err := row.Scan(
		&landmark.ID,
		&landmark.Name,
		&landmark.NativeName,
		&landmark.Description,
		&landmark.WikiURL,
		&location,
		&imgURLs,
		&landmark.Default,
		&landmark.UserID,
		&landmark.CreatedAt,
		&landmark.UpdatedAt,
		&landmark.CategoryID,
		&landmark.CreatedBy,
		&landmark.Category,
	)
	if err != nil {
		return landmark, err
	}

	landmark.Location = pgArrayToSlice(location)
	landmark.ImgURLs = pgArrayToSlice(imgURLs)

	return landmark, nil
}

// UpdateLandmark ...
func (d *Database) UpdateLandmark(formMap map[string]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE landmarks SET name=$1,native_name=$2,description=$3,wiki_url=$4,
	location=$5,img_urls=$6,category_id=$7,updated_at=$8 WHERE id=$9;
	`

	_, err := d.Conn.ExecContext(
		ctx,
		query,
		formMap["name"],
		formMap["native_name"],
		formMap["description"],
		formMap["wiki_url"],
		formMap["location"],
		formMap["img_urls"],
		formMap["category_id"],
		time.Now(),
		formMap["landmark_id"],
	)
	if err != nil {
		return err
	}

	return nil
}

// DeleteLandmark ...
func (d *Database) DeleteLandmark(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM landmarks WHERE id=$1;"

	_, err := d.Conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// SelectLandmarksWithQueries ...
func (d *Database) SelectLandmarksWithQueries(urlQueries map[string]string) ([]Landmark, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	landmarks := []Landmark{}
	landmark := Landmark{}
	location := ""
	imgURLs := ""
	baseQuery := `SELECT landmarks.*,users.username,categories.name FROM landmarks
	LEFT JOIN users ON users.id=landmarks.user_id
	LEFT JOIN categories ON categories.id=landmarks.category_id
	`

	query, err := parseLandmarkQueries(baseQuery, urlQueries)
	if err != nil {
		return landmarks, err
	}

	rows, err := d.Conn.QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		return landmarks, err
	}

	for rows.Next() {
		err := rows.Scan(
			&landmark.ID,
			&landmark.Name,
			&landmark.NativeName,
			&landmark.Description,
			&landmark.WikiURL,
			&location,
			&imgURLs,
			&landmark.Default,
			&landmark.UserID,
			&landmark.CreatedAt,
			&landmark.UpdatedAt,
			&landmark.CategoryID,
			&landmark.CreatedBy,
			&landmark.Category,
		)
		if err != nil {
			return landmarks, err
		}

		landmark.Location = pgArrayToSlice(location)
		landmark.ImgURLs = pgArrayToSlice(imgURLs)
		landmarks = append(landmarks, landmark)
	}

	return landmarks, nil
}
