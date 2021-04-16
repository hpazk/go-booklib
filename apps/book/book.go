package book

type Book struct {
	Title       string
	Description string
	Author      string
	Year        int
	IdCategory  int
	Stock       int
	Status      string
}

type Category struct {
	Name string
}
