package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // nunca expor senha no JSON
	CreatedAt time.Time `json:"created_at"`
}
