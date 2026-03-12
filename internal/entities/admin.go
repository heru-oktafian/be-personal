package entities

import "time"

type Admin struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // Tidak pernah diekspos ke JSON
	CreatedAt    time.Time `json:"created_at"`
}
