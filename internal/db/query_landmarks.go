package db

import (
	"context"
	"fmt"
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
	query := "SELECT * FROM landmarks WHERE id=$1;"
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
	query := "SELECT * FROM landmarks WHERE name=$1;"
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
	)
	if err != nil {
		return landmark, err
	}

	landmark.Location = pgArrayToSlice(location)
	landmark.ImgURLs = pgArrayToSlice(imgURLs)

	return landmark, nil
}

// InsertLandmark inserts a new landmark in the database
func (d *Database) InsertLandmark(r *http.Request, id int64, strMap map[string][]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	imgURLs, exists := strMap["img-urls"]
	if !exists {
		return fmt.Errorf("img-urls is missing")
	}

	location, exists := strMap["location"]
	if !exists {
		return fmt.Errorf("location is missing")
	}

	query := `INSERT INTO landmarks
	(name,native_name,class,description,wiki_url,
	location,img_urls,default_landmark,user_id,created_at,updated_at)
	values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);
	`
	_, err := d.Conn.ExecContext(
		ctx,
		query,
		r.PostFormValue("name"),
		r.PostFormValue("native-name"),
		r.PostFormValue("class"),
		r.PostFormValue("description"),
		r.PostFormValue("wiki-url"),
		location,
		imgURLs,
		false,
		id,
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
	query := "SELECT * FROM landmarks;"

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
