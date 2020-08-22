package models

import (
	//	"encoding/base64"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx"

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

// UserService - the interface to create, delete, and find users
type UserService interface {
	// Authenticate - authenticates an user for signing in
	Authenticate(username, password string) (*User, error)
	UserDB
}

// CreateUser - from the required fields create a User
func CreateUser(username string, email string, password string, icon string) (*User, error) {
	hashed, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		Username:  username,
		Email:     email,
		Password:  hashed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Icon:      icon,
	}, nil
}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// UserDB - the interface for interacting with DB, will be postgres for this app
type UserDB interface {
	Insert(User) (*User, error)
	Update(User) (*User, error)
	Delete(User) error
	ByUsername(User) (*User, error)
}

type userPGX struct {
	psql *pgx.Conn
}

func (u *userPGX) Insert(user User) (*User, error) {

}
