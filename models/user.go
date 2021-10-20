package models

import (
	"context"

	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

var (
	ctx = context.Background()
)

// User represents, well, a user
type User struct {
	ID       uuid.UUID `json:"-"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	// never raw, always use bcyrpt
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// will be a generated icon, the user will be able to re-generated several times a day, value will be location
	Icon string `json:"icon"`
}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
