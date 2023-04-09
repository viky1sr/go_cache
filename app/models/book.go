package models

type Book struct {
	ID     uint   `json:"id"`
	Title  string `json:"title" validate:"required,max=50"`
	Author string `json:"author" validate:"required"`
	Year   string `json:"year" validate:"required"`
}

func NewBook(title, author, year string) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Year:   year,
	}
}
