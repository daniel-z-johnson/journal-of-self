package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Create(name string, email string, pwd string) (*User, error) {
	u := &User{}
	u.Name = name
	u.Email = email
	password, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost+1)
	if err != nil {
		return nil, err
	}
	u.Password = string(password)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return u, nil
}

// timeSinceCreated
// returns the time since the user was created
func (u *User) timeSinceCreated() time.Duration {
	return time.Since(u.CreatedAt)
}
