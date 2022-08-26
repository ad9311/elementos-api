package db

import "time"

// User ...
type User struct {
	ID             int64
	FirstName      string
	LastName       string
	Username       string
	HashedPassword string
	Default        bool
	LastLogin      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Landmark ...
type Landmark struct {
	ID        int64
	UserID    int64
	Default   bool
	CreatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name        string   `json:"name"`
	NativeName  string   `json:"nativeName"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	WikiURL     string   `json:"wikiURL"`
	Location    []string `json:"location"`
	ImgURLs     []string `json:"imgURL"`
}
