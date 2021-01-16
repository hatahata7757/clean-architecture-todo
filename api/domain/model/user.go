package model

import "time"

// User は単一ユーザーの型
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Users は複数のユーザーを扱う型
type Users []User
