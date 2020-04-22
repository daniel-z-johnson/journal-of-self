package maodels

import (
	"encoding/base64"
	"time"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

// User represents, well, a user
type User struct {
	ID       uuid.UUID
	Username string
	Email    string
	// never raw, always use bcyrpt
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	// will be a generated icon, the user will be able to re-generated several times a day, value will be location
	Icon string
}

// CreateUser - from the required fields create a User
func CreateUser(username string, email string, password string, icon string) *User {
	return &User{
		Username:  username,
		Email:     email,
		Password:  hashPassword(password),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Icon:      icon,
	}
}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hashed), nil
}
