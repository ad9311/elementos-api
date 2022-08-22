package db

import "time"

// User is the users database model.
type User struct {
	ID                int64
	FirstName         string
	LastName          string
	Username          string
	Email             string
	EncryptedPassword string
	DefaultUser       bool
	LastLogin         time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// InviationCodes is the code used for sigining up.
type InviationCodes struct {
	ID        int64
	Code      string
	Validity  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Landmark ...
type Landmark struct {
	ID        int64
	CreatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name         string   `json:"name"`
	NativeName   string   `json:"nativeName"`
	Class        string   `json:"class"`
	Description  string   `json:"description"`
	StartingYear string   `json:"startingYear"`
	EndingYear   string   `json:"endingYear"`
	WikiURL      string   `json:"wikiURL"`
	Location     []string `json:"location"`
	ImgURLs      []string `json:"imgURL"`
}
