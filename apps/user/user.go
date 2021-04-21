package user

import (
	"time"

	"gorm.io/gorm"
)

// TODO 1: timestamp db and response

// User Entity
type User struct {
	gorm.Model
	Name            string
	Address         string
	Photo           string
	Email           string
	EmailVerifiedAt time.Time
	Password        string
	Role            string
}

// Request
type request struct {
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// type updateRequest struct {
// 	Name    string `json:"name" validate:"required"`
// 	Address string `json:"address" validate:"required"`
// 	Email   string `json:"email" validate:"required,email"`
// }

// Response
type response struct {
	ID              uint           `json:"id"`
	Name            string         `json:"name"`
	Address         string         `json:"address"`
	Photo           string         `json:"photo"`
	Email           string         `json:"email"`
	Role            string         `json:"role"`
	EmailVerifiedAt time.Time      `json:"email_verified_at"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
	AuthToken       string         `json:"auth_token"`
}

type loginResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Photo     string `json:"photo"`
	Email     string `json:"email"`
	AuthToken string `json:"auth_token"`
}
