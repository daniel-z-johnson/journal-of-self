package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

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
	// Authenticate(username, password string) (*User, error)
	UserDB
}

type userService struct {
	UserDB
}

func (us *userService) Insert(user User) (*User, error) {
	fmt.Println(us.UserDB)
	return us.UserDB.Insert(user)
}

func (us *userService) Update(user User) (*User, error) {
	return us.UserDB.Update(user)
}

func (us *userService) Delete(user User) error {
	return us.UserDB.Delete(user)
}

func (us *userService) ByUsername(username string) (*User, error) {
	return us.UserDB.ByUsername(username)
}

func newUserService(connection *pgxpool.Pool) UserService {
	return &userService{
		UserDB: &userPGX{
			psql: connection,
		},
	}
}

// CreateUser - from the required fields create a User
func CreateUser(username string, email string, password string, icon string) (*User, error) {
	hashed, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        uuid,
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
	ByUsername(string) (*User, error)
}

type userPGX struct {
	psql *pgxpool.Pool
}

func (u *userPGX) Insert(user User) (*User, error) {
	_, err := u.psql.Exec(ctx, `INSERT INTO users (id, username, email, password, created_at, updated_at, icon) VALUES 
													($1,      $2,     $3,       $4,        $5,         $6,   $7)`,
		user.ID, user.Username, user.Email, user.Password, time.Now(), time.Now(), user.Icon)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userPGX) Update(user User) (*User, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (u *userPGX) Delete(user User) error {
	return fmt.Errorf("Not implements")
}

func (u *userPGX) ByUsername(username string) (*User, error) {
	return nil, fmt.Errorf("Not implemented")
}
