package model

import "time"

// User structure is a type that handles a single user
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Users structure is a type that handles multiple users
type Users []User

// NewUser is a method that returns a User instance
func NewUser(id int, name string, createdAt time.Time, updatedAt time.Time) *User {
	return &User{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
