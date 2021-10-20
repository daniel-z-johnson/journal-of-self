package models

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// UserDB - the interface for interacting with DB, will be postgres for this app
type userDB interface {
	Insert(User) (*User, error)
	Update(User) (*User, error)
	Delete(User) error
	ByUsername(string) (*User, error)
}

// Make sure userPGX implements userDB
var  _ userDB = &userPGX{}

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
