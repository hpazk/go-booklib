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
	CategoryID  int
	Stock       int
	Status      string
	Category    Category
}

type Category struct {
	gorm.Model
	Name string
}

// TODO book-category relation

type request struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Author      string `json:"author" validate:"required"`
	Year        int    `json:"year" validate:"required"`
	CategoryID  int    `json:"category_id" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
	Status      string `json:"status" validate:"required"`
}

type response struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Author      string     `json:"author"`
	Year        int        `json:"year"`
	CategoryID  int        `json:"category_id"`
	Category    string     `json:"category"`
	Stock       int        `json:"stock"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
