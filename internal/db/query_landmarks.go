package db

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// SelectLandmarkByID selects a landmark by its id and returns it.
func (d *Database) SelectLandmarkByID(id int64) (*Landmark, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	landmark := Landmark{}
	query := "SELECT * FROM landmarks WHERE id=$1;"
	row := d.Conn.QueryRowContext(ctx, query, id)
	err := row.Scan()
	if err != nil {
		return &landmark, err
	}

	return &landmark, nil
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
	location,img_urls,users_id,created_at,updated_at)
	values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);
	`
	res, err := d.Conn.ExecContext(
		ctx,
		query,
		r.PostFormValue("name"),
		r.PostFormValue("native-name"),
		r.PostFormValue("class"),
		r.PostFormValue("description"),
		r.PostFormValue("wiki-url"),
		location,
		imgURLs,
		id,
		time.Now(),
		time.Now(),
	)
	fmt.Println(res)

	if err != nil {
		return err
	}

	return nil
}
