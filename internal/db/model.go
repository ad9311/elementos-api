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
