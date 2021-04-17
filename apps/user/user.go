package user

import (
	"time"

	"github.com/jinzhu/gorm"
)

// TODO 1: gorm.Model

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

type updateRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}

// Response
type response struct {
	ID              uint       `json:"id"`
	Name            string     `json:"name"`
	Address         string     `json:"address"`
	Photo           string     `json:"photo"`
	Email           string     `json:"email"`
	EmailVerifiedAt time.Time  `json:"email_verified_at"`
	AuthToken       string     `json:"auth_token"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}

func userResponseFormatter(user User, auth_token string) response {
	formatter := response{
		ID:              user.ID,
		Name:            user.Name,
		Address:         user.Address,
		Photo:           user.Photo,
		Email:           user.Email,
		EmailVerifiedAt: user.EmailVerifiedAt,
		AuthToken:       auth_token,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		DeletedAt:       user.DeletedAt,
	}

	return formatter
}
