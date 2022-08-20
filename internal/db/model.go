package db

import "time"

// User is the users database model.
type User struct {
	ID                int64     `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"password"`
	DefaultUser       bool      `json:"default_user"`
	LastLogin         time.Time `json:"last_login"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// InviationCodes is the code used for sigining up.
type InviationCodes struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Validity  time.Time `json:"validity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
