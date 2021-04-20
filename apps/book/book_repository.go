package book

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type bookRepository interface {
	Store(bokk Book) (Book, error)
	Fetch() ([]Book, error)
	FetchByCategory(category string) ([]Book, error)
	FindCategory(categoryName string) (Category, error)
	FindByNewest() ([]Book, error)
	// Update(book Book) (Book, error)
	// Delete(id uint) error
	// FindById(id uint) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func BookRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Save New Book
func (r *repository) Store(user Book) (Book, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Get All Book
func (r *repository) Fetch() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}

func (r *repository) FindCategory(categoryName string) (Category, error) {
	var category Category
	err := r.db.Where("name = ?", categoryName).First(&category).Error
	if err != nil {
		return category, err
	}

	fmt.Println(category)

	return category, nil
}

func (r *repository) FindByNewest() ([]Book, error) {
	var books []Book
	err := r.db.Where("year >= ?", time.Now().Year()-2).Limit(20).Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}

func (r *repository) FetchByCategory(category string) ([]Book, error) {
	var books []Book

	err := r.db.Raw("SELECT * FROM books WHERE category_id = (SELECT id FROM categories WHERE name = ?)", category).Scan(&books).Error
	if err != nil {
		return books, nil
	}

	fmt.Println(len(books))
	fmt.Println(category)

	return books, nil
}

// Get User by Id
func (r *repository) FindById(id uint) (Book, error) {
	var book Book

	err := r.db.First(&book, "id = ?", id).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

// Update user
func (r *repository) Update(book Book) (Book, error) {
	err := r.db.Save(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

// // Delete User
// func (r *repository) Delete(id uint) error {
// 	usersStorage[id-1] = usersStorage[len(usersStorage)-1]
// 	usersStorage[uint(len(usersStorage))-1] = User{}
// 	usersStorage = usersStorage[:uint(len(usersStorage))-1]
// 	return nil
// }
