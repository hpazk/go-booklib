package book

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Description string
	Author      string
	Year        int
	IdCategory  int
	Stock       int
	Status      string
}

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

// TODO book-category relation

type request struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Author      string `json:"author" validate:"required"`
	Year        int    `json:"year" validate:"required"`
	IdCategory  int    `json:"id_category" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
	Status      string `json:"status" validate:"required"`
}

type response struct {
	ID              uint       `json:"id"`
	Title           string     `json:"title" validate:"required"`
	Description     string     `json:"description" validate:"required"`
	Author          string     `json:"author" validate:"required"`
	Year            int        `json:"year" validate:"required"`
	IdCategory      int        `json:"id_category" validate:"required"`
	Stock           int        `json:"stock" validate:"required"`
	Status          string     `json:"status" validate:"required"`
	EmailVerifiedAt time.Time  `json:"email_verified_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}
