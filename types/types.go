// Package types contains all different types that's gonna be used all over the project
package types

import "time"

type RegisterUserPayload struct {
	FirstName string `json:"fistName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"fistName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
