package book

func bookResponseFormatter(book Book) response {
	formatter := response{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
		Year:        book.Year,
		CategoryID:  book.CategoryID,
		Category:    book.Category.Name,
		Stock:       book.Stock,
		Status:      book.Status,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
		DeletedAt:   book.DeletedAt,
	}

	return formatter
}

func booksResponseFormatter(books []Book) []response {
	formatter := []response{}

	for _, campaign := range books {
		c := bookResponseFormatter(campaign)
		formatter = append(formatter, c)
	}
	return formatter
}
