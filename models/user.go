package models

import (
	"context"

	//	"encoding/base64"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

var (
	ctx = context.Background()
)

// User represents, well, a user
type User struct {
	ID       uuid.UUID
	Username string `json:"username"`
	Email    string `json:"email"`
	// never raw, always use bcyrpt
	Password  string `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// will be a generated icon, the user will be able to re-generated several times a day, value will be location
	Icon string
}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
