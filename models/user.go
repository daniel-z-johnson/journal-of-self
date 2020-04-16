package maodels

import (
	"time"

	"github.com/google/uuid"
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
