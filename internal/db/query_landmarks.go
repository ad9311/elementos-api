package db

import (
	"context"
	"net/http"
	"time"
)

// SelectLandmarkByID selects a landmark by its id and returns it.
func (d *Database) SelectLandmarkByID(id int64) (Landmark, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	landmark := Landmark{}
	location := ""
	imgURLs := ""
	query := `SELECT landmarks.*,users.username
	FROM landmarks RIGHT JOIN users ON users.id=landmarks.user_id WHERE landmarks.id=$1`

	row := d.Conn.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&landmark.ID,
		&landmark.Name,
		&landmark.NativeName,
		&landmark.Class,
		&landmark.Description,
		&landmark.WikiURL,
		&location,
		&imgURLs,
		&landmark.Default,
		&landmark.UserID,
		&landmark.CreatedAt,
		&landmark.UpdatedAt,
		&landmark.CreatedBy,
	)
	if err != nil {
		return landmark, err
	}

	landmark.Location = pgArrayToSlice(location)
	landmark.ImgURLs = pgArrayToSlice(imgURLs)

	return landmark, nil
}

// SelectLandmarkByName selects a landmark by its name and returns it.
func (d *Database) SelectLandmarkByName(name string) (Landmark, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	landmark := Landmark{}
	location := ""
	imgURLs := ""
	query := `SELECT landmarks.*,users.username
	FROM landmarks RIGHT JOIN users ON users.id=landmarks.user_id WHERE landmarks.name=$1`

	row := d.Conn.QueryRowContext(ctx, query, name)
	err := row.Scan(
		&landmark.ID,
		&landmark.Name,
		&landmark.NativeName,
		&landmark.Class,
		&landmark.Description,
		&landmark.WikiURL,
		&location,
		&imgURLs,
		&landmark.Default,
		&landmark.UserID,
		&landmark.CreatedAt,
		&landmark.UpdatedAt,
		&landmark.CreatedBy,
	)
	if err != nil {
		return landmark, err
	}

	landmark.Location = pgArrayToSlice(location)
	landmark.ImgURLs = pgArrayToSlice(imgURLs)

	return landmark, nil
}

// InsertLandmark inserts a new landmark in the database
func (d *Database) InsertLandmark(r *http.Request, formMap map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO landmarks
	(name,native_name,class,description,wiki_url,
	location,img_urls,default_landmark,user_id,created_at,updated_at)
	values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);
	`
	_, err := d.Conn.ExecContext(
		ctx,
		query,
		formMap["name"],
		formMap["native-name"],
		formMap["class"],
		formMap["description"],
		formMap["wiki-url"],
		formMap["location"],
		formMap["img-urls"],
		false,
		formMap["user-id"],
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

// SelectLandmarks selects all landmarks from the database.
func (d *Database) SelectLandmarks() ([]*Landmark, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	landmarks := []*Landmark{}
	landmark := Landmark{}
	location := ""
	imgURLs := ""
	query := `SELECT landmarks.*,users.username
	FROM landmarks RIGHT JOIN users ON users.id=landmarks.user_id`

	rows, err := d.Conn.QueryContext(ctx, query)
	if err != nil {
		rows.Close()
		return landmarks, err
	}

	for rows.Next() {
		err := rows.Scan(&landmark.ID,
			&landmark.Name,
			&landmark.NativeName,
			&landmark.Class,
			&landmark.Description,
			&landmark.WikiURL,
			&location,
			&imgURLs,
			&landmark.Default,
			&landmark.UserID,
			&landmark.CreatedAt,
			&landmark.UpdatedAt,
			&landmark.CreatedBy,
		)
		if err != nil {
			rows.Close()
			return landmarks, err
		}

		lm := landmark
		lm.Location = pgArrayToSlice(location)
		lm.ImgURLs = pgArrayToSlice(imgURLs)
		landmarks = append(landmarks, &lm)
	}
	rows.Close()

	return landmarks, nil
}

// UpdateLandmarkByID ...
func (d *Database) UpdateLandmarkByID(formMap map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE landmarks SET name=$1,native_name=$2,class=$3,description=$4,wiki_url=$5,
	location=$6,img_urls=$7,updated_at=$8 WHERE id=$9`

	_, err := d.Conn.ExecContext(
		ctx,
		query,
		formMap["name"],
		formMap["native-name"],
		formMap["class"],
		formMap["description"],
		formMap["wiki-url"],
		formMap["location"],
		formMap["img-urls"],
		time.Now(),
		formMap["landmark-id"],
	)
	if err != nil {
		return err
	}

	return nil
}
