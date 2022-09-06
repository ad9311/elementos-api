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

// Inviation ...
type Inviation struct {
	ID        int64
	Code      string
	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Landmark ...
type Landmark struct {
	ID          int64     `json:"-"`
	UserID      int64     `json:"-"`
	CategoryID  int64     `json:"-"`
	Default     bool      `json:"-"`
	CreatedBy   string    `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Name        string    `json:"name"`
	NativeName  string    `json:"nativeName"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	WikiURL     string    `json:"wikiURL"`
	Location    []string  `json:"location"`
	ImgURLs     []string  `json:"imgURL"`
}

// ResponseWithData ...
type ResponseWithData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Response ...
type Response struct {
	Message string `json:"message"`
}

// Category ...
type Category struct {
	ID        int64     `json:"-"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
