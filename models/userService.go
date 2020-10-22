package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

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
