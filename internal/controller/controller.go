package controller

import (
	"github.com/ad9311/hitomgr/internal/db"
	"github.com/ad9311/hitomgr/internal/sess"
	"golang.org/x/crypto/bcrypt"
)

// Data ...
var Data *sess.Data
var database *db.Database

// Init ...
func Init(dtbs *db.Database, sessionData *sess.Data) {
	database = dtbs
	Data = sessionData
}

func encryptPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return string(encryptedPassword), err
	}

	return string(encryptedPassword), nil
}
