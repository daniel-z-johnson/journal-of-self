package models

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UsersDB interface {
	Create(name string, email string, pwd string) (*User, error)
	Update(id int64, name string, email string, pwd string) (*User, error)
	GetByID(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByName(name string) (*User, error)
	Delete(id int64) error
}

type usersDB struct {
	db *pgxpool.Conn
}

func NewUsersDB(db *pgxpool.Conn) UsersDB {
	return &usersDB{db: db}
}

func (u *usersDB) Create(name string, email string, pwd string) (*User, error) {
	user, err := Create(name, email, pwd)
	if err != nil {
		return nil, err
	}
	err = u.db.QueryRow(context.Background(), `
		INSERT INTO users (name, email, password) VALUES
		($1, $2, $3) RETURNING id
	`, user.Name, user.Email, user.Password).Scan(user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usersDB) Update(id int64, name string, email string, pwd string) (*User, error) {
	panic("not implemented")
}

func (u *usersDB) GetByID(id int64) (*User, error) {
	panic("not implemented")
}

func (u *usersDB) GetByEmail(email string) (*User, error) {
	panic("not implemented")
}

func (u *usersDB) GetByName(name string) (*User, error) {
	panic("not implemented")
}

func (u *usersDB) Delete(id int64) error {
	panic("not implemented")
}
