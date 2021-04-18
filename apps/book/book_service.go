package book

import (
	"fmt"
	"time"
)

type BookServices interface {
	AddBook(req *request) (Book, error)
	FetchBooks() ([]Book, error)
	FetchByCategory(categoryName string) (uint, error)
	FetchNewestBooks() ([]Book, error)
	FetchBooksByCategory(categoryId uint) ([]Book, error)
}

type services struct {
	repository bookRepository
}

func BookService(repository bookRepository) *services {
	return &services{repository}
}

func (s *services) AddBook(req *request) (Book, error) {
	book := Book{}
	book.Title = req.Title
	book.Description = req.Description
	book.Author = req.Author
	book.Year = req.Year
	book.CategoryId = req.CategoryId
	book.Stock = req.Stock
	book.Status = req.Status
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	newBook, err := s.repository.Store(book)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (s *services) FetchBooks() ([]Book, error) {
	var books []Book

	books, err := s.repository.Fetch()
	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *services) FetchNewestBooks() ([]Book, error) {
	var books []Book

	books, err := s.repository.FindByNewest()
	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *services) FetchBooksByCategory(categoryId uint) ([]Book, error) {
	var books []Book

	books, err := s.repository.FetchByCategory(categoryId)
	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *services) FetchByCategory(categoryName string) (uint, error) {
	var category Category

	category, err := s.repository.FindCategory(categoryName)
	if err != nil {
		return category.ID, err
	}

	fmt.Println(category.ID)

	return category.ID, nil
}
