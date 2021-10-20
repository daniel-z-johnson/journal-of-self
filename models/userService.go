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
	userDB
}

type userService struct {
	userDB
}

func (us *userService) Insert(user User) (*User, error) {
	fmt.Println(us.userDB)
	return us.userDB.Insert(user)
}

func (us *userService) Update(user User) (*User, error) {
	return us.userDB.Update(user)
}

func (us *userService) Delete(user User) error {
	return us.userDB.Delete(user)
}

func (us *userService) ByUsername(username string) (*User, error) {
	return us.userDB.ByUsername(username)
}

func newUserService(connection *pgxpool.Pool) UserService {
	return &userService{
		userDB: &userPGX{
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
