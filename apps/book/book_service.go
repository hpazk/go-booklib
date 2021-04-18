package book

import (
	"time"
)

type BookServices interface {
	AddBook(req *request) (Book, error)
	FetchBooks() ([]Book, error)
	FetchNewestBooks() ([]Book, error)
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
	book.IdCategory = req.IdCategory
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

	books, err := s.repository.FetchByNewest()
	if err != nil {
		return books, err
	}

	return books, nil
}
